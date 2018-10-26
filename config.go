package controller

import (
    "fmt"
    "encoding/xml"
    "os"
    "io/ioutil"
    "path/filepath"
)

type PhaseTime struct {
    total_seconds int
    min_seconds int
    ending_seconds int
}

type Phase struct {
    id uint16
    name string
    time PhaseTime
    next_phase *Phase
}

// XML structures

type XmlPhaseTime struct {
    XMLName xml.Name `xml:"time"`
    TotalSeconds int `xml:"total_seconds"`
    MinSeconds int `xml:"min_seconds"`
    EndingSeconds int `xml:"ending_seconds"`
}

type XmlPhase struct {
    XMLName xml.Name `xml:"phase"`
    Id uint16 `xml:"id"`
    Name string `xml:"name"`
    Time XmlPhaseTime `xml:"time"`
    NextPhase int `xml:"next_phase"`
}

type XmlPhases struct {
    XMLName xml.Name `xml:"phases"`
    XmlPhases []XmlPhase `xml:"phase"`
}

// Load phases from XML-config (data/controller.xml)

func loadPhases(configfile string) XmlPhases {

    phasesPath, err := filepath.Abs(configfile)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    file, err := os.Open(phasesPath)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    byteValue, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    var xmlPhases XmlPhases
    xml.Unmarshal(byteValue, &xmlPhases)
    return xmlPhases
}
