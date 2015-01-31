/*
* Archon Login Server
* Copyright (C) 2014 Andrew Rodman
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program.  If not, see <http://www.gnu.org/licenses/>.
* ---------------------------------------------------------------------
*
* Singleton package for handling the login and character server configuration. Also
* responsible for establishing a connection to the database to be maintained
* during execution.
 */
package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"libarchon/util"
	"os"
	"strconv"
	"strings"
	"time"
)

const loginConfigFile = "login_config.json"

type LogType byte
type LogPriority byte

// Constants for the configurable log level that control the amount of information
// written to the server logs. The higher the number, the lower the priority.
const (
	LogPriorityCritical = 1
	LogPriorityHigh     = 2
	LogPriorityMedium   = 3
	LogPriorityLow      = 4

	LogTypeInfo    = 1 << iota
	LogTypeWarning = 1 << iota
	LogTypeError   = 1 << iota
)

// Logs a message to either the user's configured logfile or to standard out. Only messages
// equal to or greater than the user's specified priority will be written.
func LogMsg(message string, logType LogType, priority LogPriority) {
	config := GetConfig()
	if priority > config.LogLevel {
		return
	}
	var logMsg string
	timestamp := time.Now().Format("06-01-02 15:04:05")
	switch logType {
	case LogTypeError:
		logMsg = fmt.Sprintf("%s [ERROR] %s\n", timestamp, message)
	case LogTypeWarning:
		logMsg = fmt.Sprintf("%s [WARNING] %s\n", timestamp, message)
	case LogTypeInfo:
		logMsg = fmt.Sprintf("%s [INFO] %s\n", timestamp, message)
	}
	if config.Logfile == "" {
		fmt.Printf(logMsg)
	} else {
		logfile, err := os.OpenFile("login_server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("\nWARNING: Failed to open log file %s: %s\n", logfile, err.Error())
			return
		}
		defer logfile.Close()
		_, err = logfile.WriteString(logMsg)
		if err != nil {
			fmt.Printf("\nWARNING: Error writing to log file: %s\n", err)
		}
	}
}

type configuration struct {
	Hostname      string
	LoginPort     string
	CharacterPort string
	DBHost        string
	DBPort        string
	DBName        string
	DBUsername    string
	DBPassword    string
	Logfile       string
	LogLevel      LogPriority
	DebugMode     bool

	database *sql.DB
}

// Singleton instance.
var loginConfig *configuration = nil
var cachedHostBytes [4]byte

// This functiion should be used to get access to the server config instead of directly
// referencing the loginConfig pointer.
func GetConfig() *configuration {
	if loginConfig == nil {
		loginConfig = new(configuration)
	}
	return loginConfig
}

// Populate config with the contents of a JSON file at path fileName. Config parameters
// in the file must match the above fields exactly in order to be read.
func (config *configuration) InitFromFile(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	json.Unmarshal(data, config)
	config.enforceDefaults()
	return nil
}

// Provide default values for fields that are optional or critical.
func (config *configuration) enforceDefaults() {
	if config.Hostname == "" {
		config.Hostname = "127.0.0.1"
	}
	if config.LoginPort == "" {
		config.LoginPort = "12000"
	}
	if config.CharacterPort == "" {
		config.CharacterPort = "12001"
	}
	if config.LogLevel < LogPriorityCritical || config.LogLevel > LogPriorityLow {
		// The log level must be at least open to critical messages.
		config.LogLevel = LogPriorityCritical
	}
}

// Convert the hostname string into 4 bytes to be used with the redirect packet.
func (config *configuration) HostnameBytes() [4]byte {
	// Hacky, but chances are the IP address isn't going to start with 0 and a
	// fixed-length array can't be null.
	if cachedHostBytes[0] == 0x00 {
		parts := strings.Split(config.Hostname, ".")
		for i := 0; i < 4; i++ {
			tmp, _ := strconv.ParseUint(parts[i], 10, 8)
			cachedHostBytes[i] = uint8(tmp)
		}
	}
	return cachedHostBytes
}

// Establish a connection to the database and ping it to verify.
func (config *configuration) InitDb() error {
	dbName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUsername,
		config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	var err error
	config.database, err = sql.Open("mysql", dbName)
	if err != nil || config.database.Ping() != nil {
		return err
	}
	return nil
}

func (config *configuration) CloseDb() {
	config.database.Close()
}

// Use this function to obtain a reference to the database so that it can remain
// encapsulated and any consistency checks can be centralized.
func (config *configuration) Database() *sql.DB {
	if config.database == nil {
		// Don't implicitly initialize the database - if there's an error or other action that causes
		// the reference to become nil then we're probably leaking a connection.
		panic(util.ServerError{Message: "Attempt to reference uninitialized database"})
	}
	return config.database
}

func (config *configuration) String() string {
	logfile := config.Logfile
	if logfile == "" {
		logfile = "Standard Out"
	}
	return "Hostname: " + config.Hostname + "\n" +
		"Login Port: " + config.LoginPort + "\n" +
		"Character Port: " + config.CharacterPort + "\n" +
		"Database Host: " + config.DBHost + "\n" +
		"Database Port: " + config.DBPort + "\n" +
		"Database Name: " + config.DBName + "\n" +
		"Database Username: " + config.DBUsername + "\n" +
		"Database Password: " + config.DBPassword + "\n" +
		"Output Logged To: " + logfile + "\n" +
		"Logging Level: " + strconv.FormatInt(int64(config.LogLevel), 10) + "\n" +
		"Debug Mode Enabled: " + strconv.FormatBool(config.DebugMode)
}
