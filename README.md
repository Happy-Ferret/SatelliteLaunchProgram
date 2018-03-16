## Problem
Let's say, Indian Space Research Organization(ISRO) wants to launch 500 
satellites on a certain date. 
The launch of this will happen from 2 separate Launch Sites(LS).
Lets call them LS1 and LS2.
The way the Satellites(SATs) are split for each of the locations is:
SAT1, SAT2, SAT3 and SAT4 are going to be launched from LS1. 
SAT5, SAT6, SAT7 and SAT8 are going to launch from LS2.
Subsequently SAT9, SAT10 SAT11 and SAT12 will be launched from LS1 
and SAT13, SAT14 SAT15 and SAT16 launched from LS2. This process goes on until all the 500 satellites are launched. 

There are a couple of constraints that ISRO needs to work under:
1. The group of 4 satellites that need to be launched from each Launch Site needs to happen simultaneously 
  with each satellite counting down and launching.
2. The launch of the Satellite group happens in tandem. What it means is, SATs 1-4 launched from LS1 needs
  to notify LS2 of the successful launch and SATs 5-8 then are launched from  LS2. The same drill happens, 
  LS2 notifies LS1 of the launch and LS1 would launch SATs 9-12. This happens until till the time all the satellites are launched.

 Please build a GOLang program to get make the launch and the communication happen. 
 Please upload the program on Github and send me the instructions to checkout the code and run it on our machines.

## Download and run

Locally
```
git clone https://github.com/JessTheBell/SatelliteLaunchProgram
cd SatelliteLaunchProgram
go build main.go
./main
```

On the Go Playground

[Open In Go Playground](https://play.golang.org/p/R2lR1NMVPFQ)
