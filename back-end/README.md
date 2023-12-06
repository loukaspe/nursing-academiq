# Nursing Academiq REST API

---

## Description

Service that provides a REST API offering CRUD operations for User, Student, Tutor

---

## Run

`make start-app`

* This command will start the app with `localhost` address and `:8080` port (specified in Dockerfile.dev and .env)

---

## Makefile Commands

| Command                         | Usage                                                                  |
|---------------------------------|------------------------------------------------------------------------|
| start-app                       | `Start app`                                                            |
| kill-app                        | `Stop app`                                                             |
| rebuild-app                     | `Rebuild app`                                                          |
| tests-all                       | `Run both unit and integration tests`                                  |
| tests-benchmark                 | `Run benchmark tests`                                                  |
| tests-unit                      | `Run unit tests `                                                      |
| tests-file FILE={filePath}      | `Run specific file test`                                               |
| generate-mock FILE={filePath}   | `Generate mock for a specific file`                                    |
| tests-package PACKAGE={package} | `Run specific package test`                                            |
| tests-all-with-coverage         | `Run both unit and integration tests via docker with coverage details` |

* All these are executed through docker containers
* In order to execute makefile commands type **make** plus a command from the table above

  make {command}

---

## Notes

1. There are three Dockerfile files.
    1. Dockerfile is the normal, production one
    2. Dockerfile.dev is for setting up a remote debugger Delve
    3. utilities.Dockerfile is for building a docker for "utilities" like running tests,  
       linting etc