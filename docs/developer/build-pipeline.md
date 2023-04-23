# Build Pipeline

## Platform

http://drone.akkadius.com/EQEmu/Server

The current CI pipeline uses [Drone](https://www.drone.io/) which is written in Golang and [open source](https://github.com/harness/drone)

![image](https://user-images.githubusercontent.com/3319450/214195643-43bb8b8b-04ca-42c7-9cf0-9bf612dc2016.png)

## Configuration

Build configuration is stored `.drone.yml` at the top level of the codebase. Drone configuration reference docs can be found [here](https://docs.drone.io/pipeline/overview/) if a future developer finds themselves navigating it. Is is unlikely and very infrequent that someone should need to

Build scripts are located in `./utils/scripts/build/`

## Hardware

The build pipeline runs on dedicated hardware to keep feedback loops tight for developers

| **Build Type**   | **CPU**  | **RAM**  | **Disk**     |  
|:---: |:---: |:---: |:---: |  
| Windows  | Ryzen 9 5900HX (8C/16T, Max 4.6Ghz)  | 64GB DDR4    | WD_BLACK 1TB 7,300 MB/s  |  
| Linux    | Ryzen 9 5900HX (8C/16T, Max 4.6Ghz)  | 64GB DDR4    | WD_BLACK 1TB 7,300 MB/s  |

## Time

Build times take roughly 3-5 minutes on average. Both Linux and Windows running in parallel in their own steps

![image](https://user-images.githubusercontent.com/3319450/214195050-0581abba-13c3-4824-8e96-68b899b86408.png)
