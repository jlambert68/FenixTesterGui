package UnitTestTestData

import (
	fenixTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

// Availalbe Buoilding Blocks; ABB

// ******* START ABB001 *******
// gRPC-result when asking for Available Building Blocks
var TestInstructionsAndTestInstructionsContainersRespons_ABB001 string = `{
  "ImmatureTestInstructions": [
    {
      "BasicTestInstructionInformation": {
        "NonEditableInformation": {
          "DomainUuid": "78a97c41-a098-4122-88d2-01ed4b6c4844",
          "DomainName": "Custody Arrangement",
          "TestInstructionUuid": "2f130d7e-f8aa-466f-b29d-0fb63608c1a6",
          "TestInstructionName": "Just the name",
          "TestInstructionTypeUuid": "513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb",
          "TestInstructionTypeName": "The type of the TestInstruction",
          "Deprecated": false,
          "MajorVersionNumber": 1,
          "MinorVersionNumber": 0,
          "UpdatedTimeStamp": "2022-06-19T14:11:52Z",
          "TestInstructionColor": "#00ff00",
          "TCRuleDeletion": "TCRuleDeletion020",
          "TCRuleSwap": "TCRuleSwap020"
        },
        "EditableInformation": {
          "TestInstructionDescription": "A TestIntruction that's cool",
          "TestInstructionMouseOverText": "This will be shown when hovering above this TestInstruction"
        },
        "InvisibleBasicInformation": {
          "Enabled": true
        }
      },
      "ImmatureTestInstructionInformation": {
        "AvailableDropZones": [
          {
            "DropZoneUuid": "a5c27024-e40c-49f7-8667-eab485c65105",
            "DropZoneName": "My First DropZone",
            "DropZoneDescription": "This is the Description of my first DropZone",
            "DropZoneMouseOver": "This is the mouseover text for my first DropZone",
            "DropZoneColor": "#0000ff",
            "DropZonePreSetTestInstructionAttributes": [
              {
                "TestInstructionAttributeType": "TEXTBOX",
                "TestInstructionAttributeUuid": "7fb2566e-bf27-44fd-818c-4591fa9c603c",
                "TestInstructionAttributeName": "This attributes holds a Text",
                "AttributeValueAsString": "Jonas",
                "AttributeValueUuid": "7fb2566e-bf27-44fd-818c-4591fa9c603c"
              }
            ]
          }
        ]
      },
      "ImmatureSubTestCaseModel": {
        "FirstImmatureElementUuid": "",
        "TestCaseModelElements": [
          {
            "OriginalElementUuid": "2f130d7e-f8aa-466f-b29d-0fb63608c1a6",
            "OriginalElementName": "Just the name",
            "MatureElementUuid": "",
            "PreviousElementUuid": "2f130d7e-f8aa-466f-b29d-0fb63608c1a6",
            "NextElementUuid": "2f130d7e-f8aa-466f-b29d-0fb63608c1a6",
            "FirstChildElementUuid": "2f130d7e-f8aa-466f-b29d-0fb63608c1a6",
            "ParentElementUuid": "2f130d7e-f8aa-466f-b29d-0fb63608c1a6",
            "TestCaseModelElementType": "TI_TESTINSTRUCTION",
            "CurrentElementModelElement": "2f130d7e-f8aa-466f-b29d-0fb63608c1a6"
          }
        ]
      }
    }
  ],
  "ImmatureTestInstructionContainers": [
    {
      "BasicTestInstructionContainerInformation": {
        "NonEditableInformation": {
          "DomainUuid": "78a97c41-a098-4122-88d2-01ed4b6c4844",
          "DomainName": "Custody Arrangement",
          "TestInstructionContainerUuid": "aa1b9734-5dce-43c9-8d77-3368940cf126",
          "TestInstructionContainerName": "Emtpy parallelled processed Turbo TestInstructionsContainer",
          "TestInstructionContainerTypeUuid": "ca07bdd9-4152-4020-b3f0-fc750b45885e",
          "TestInstructionContainerTypeName": "CA Base containers",
          "Deprecated": false,
          "MajorVersionNumber": 1,
          "MinorVersionNumber": 0,
          "UpdatedTimeStamp": "2022-06-22T17:15:16Z",
          "TestInstructionContainerColor": "#aaaaaa",
          "TCRuleDeletion": "TCRuleDeletion011",
          "TCRuleSwap": "TCRuleSwap012"
        },
        "EditableInformation": {
          "TestInstructionContainerDescription": "Children of this CA-container is processed in parallel",
          "TestInstructionContainerMouseOverText": "Children of this CA-container is processed in parallel"
        },
        "InvisibleBasicInformation": {
          "Enabled": true
        },
        "EditableTestInstructionContainerAttributes": {
          "TestInstructionContainerExecutionType": "PARALLELLED_PROCESSED"
        }
      },
      "ImmatureTestInstructionContainerInformation": {
        "AvailableDropZones": []
      },
      "ImmatureSubTestCaseModel": {
        "FirstImmatureElementUuid": "",
        "TestCaseModelElements": [
          {
            "OriginalElementUuid": "aa1b9734-5dce-43c9-8d77-3368940cf126",
            "OriginalElementName": "Emtpy parallelled processed Turbo TestInstructionsContainer",
            "MatureElementUuid": "",
            "PreviousElementUuid": "0883d538-1cff-4be1-ba1f-4dc1f68f6242",
            "NextElementUuid": "0883d538-1cff-4be1-ba1f-4dc1f68f6242",
            "FirstChildElementUuid": "0883d538-1cff-4be1-ba1f-4dc1f68f6242",
            "ParentElementUuid": "aa1b9734-5dce-43c9-8d77-3368940cf126",
            "TestCaseModelElementType": "B10_BOND",
            "CurrentElementModelElement": "0883d538-1cff-4be1-ba1f-4dc1f68f6242"
          },
          {
            "OriginalElementUuid": "aa1b9734-5dce-43c9-8d77-3368940cf126",
            "OriginalElementName": "Emtpy parallelled processed Turbo TestInstructionsContainer",
            "MatureElementUuid": "",
            "PreviousElementUuid": "aa1b9734-5dce-43c9-8d77-3368940cf126",
            "NextElementUuid": "aa1b9734-5dce-43c9-8d77-3368940cf126",
            "FirstChildElementUuid": "0883d538-1cff-4be1-ba1f-4dc1f68f6242",
            "ParentElementUuid": "aa1b9734-5dce-43c9-8d77-3368940cf126",
            "TestCaseModelElementType": "TIC_TESTINSTRUCTIONCONTAINER",
            "CurrentElementModelElement": "aa1b9734-5dce-43c9-8d77-3368940cf126"
          }
        ]
      }
    },
    {
      "BasicTestInstructionContainerInformation": {
        "NonEditableInformation": {
          "DomainUuid": "e81b9734-5dce-43c9-8d77-3368940cf126",
          "DomainName": "Fenix",
          "TestInstructionContainerUuid": "e107bdd9-4152-4020-b3f0-fc750b45885e",
          "TestInstructionContainerName": "Emtpy parallelled processed TestInstructionsContainer",
          "TestInstructionContainerTypeUuid": "b107bdd9-4152-4020-b3f0-fc750b45885e",
          "TestInstructionContainerTypeName": "Base containers",
          "Deprecated": false,
          "MajorVersionNumber": 1,
          "MinorVersionNumber": 0,
          "UpdatedTimeStamp": "2022-06-22T17:15:16Z",
          "TestInstructionContainerColor": "#aaaaaa",
          "TCRuleDeletion": "TCRuleDeletion011",
          "TCRuleSwap": "TCRuleSwap012"
        },
        "EditableInformation": {
          "TestInstructionContainerDescription": "Children of this container is processed in parallel",
          "TestInstructionContainerMouseOverText": "Children of this container is processed in parallel"
        },
        "InvisibleBasicInformation": {
          "Enabled": true
        },
        "EditableTestInstructionContainerAttributes": {
          "TestInstructionContainerExecutionType": "PARALLELLED_PROCESSED"
        }
      },
      "ImmatureTestInstructionContainerInformation": {
        "AvailableDropZones": []
      },
      "ImmatureSubTestCaseModel": {
        "FirstImmatureElementUuid": "",
        "TestCaseModelElements": [
          {
            "OriginalElementUuid": "e107bdd9-4152-4020-b3f0-fc750b45885e",
            "OriginalElementName": "Emtpy parallelled processed TestInstructionsContainer",
            "MatureElementUuid": "",
            "PreviousElementUuid": "0883d538-1cff-4be1-ba1f-4dc1f68f6242",
            "NextElementUuid": "0883d538-1cff-4be1-ba1f-4dc1f68f6242",
            "FirstChildElementUuid": "0883d538-1cff-4be1-ba1f-4dc1f68f6242",
            "ParentElementUuid": "e107bdd9-4152-4020-b3f0-fc750b45885e",
            "TestCaseModelElementType": "B10_BOND",
            "CurrentElementModelElement": "0883d538-1cff-4be1-ba1f-4dc1f68f6242"
          },
          {
            "OriginalElementUuid": "e107bdd9-4152-4020-b3f0-fc750b45885e",
            "OriginalElementName": "Emtpy parallelled processed TestInstructionsContainer",
            "MatureElementUuid": "",
            "PreviousElementUuid": "e107bdd9-4152-4020-b3f0-fc750b45885e",
            "NextElementUuid": "e107bdd9-4152-4020-b3f0-fc750b45885e",
            "FirstChildElementUuid": "0883d538-1cff-4be1-ba1f-4dc1f68f6242",
            "ParentElementUuid": "e107bdd9-4152-4020-b3f0-fc750b45885e",
            "TestCaseModelElementType": "TIC_TESTINSTRUCTIONCONTAINER",
            "CurrentElementModelElement": "e107bdd9-4152-4020-b3f0-fc750b45885e"
          }
        ]
      }
    },
    {
      "BasicTestInstructionContainerInformation": {
        "NonEditableInformation": {
          "DomainUuid": "e81b9734-5dce-43c9-8d77-3368940cf126",
          "DomainName": "Fenix",
          "TestInstructionContainerUuid": "f81b9734-5dce-43c9-8d77-3368940cf126",
          "TestInstructionContainerName": "Emtpy serial processed TestInstructionsContainer",
          "TestInstructionContainerTypeUuid": "b107bdd9-4152-4020-b3f0-fc750b45885e",
          "TestInstructionContainerTypeName": "Base containers",
          "Deprecated": false,
          "MajorVersionNumber": 1,
          "MinorVersionNumber": 0,
          "UpdatedTimeStamp": "2022-06-22T17:15:16Z",
          "TestInstructionContainerColor": "#aaaaaa",
          "TCRuleDeletion": "TCRuleDeletion011",
          "TCRuleSwap": "TCRuleSwap012"
        },
        "EditableInformation": {
          "TestInstructionContainerDescription": "Children of this container is processed in serial",
          "TestInstructionContainerMouseOverText": "Children of this container is processed in serial"
        },
        "InvisibleBasicInformation": {
          "Enabled": true
        },
        "EditableTestInstructionContainerAttributes": {
          "TestInstructionContainerExecutionType": "SERIAL_PROCESSED"
        }
      },
      "ImmatureTestInstructionContainerInformation": {
        "AvailableDropZones": [
          {
            "DropZoneUuid": "c5e37024-e40c-49f7-8667-eab485c65105",
            "DropZoneName": "My first DropZone for a TestInstructionContainer",
            "DropZoneDescription": "This is the Description of my first DropZone",
            "DropZoneMouseOver": "This is the mouseover text for my first DropZone",
            "DropZoneColor": "#0000e",
            "DropZonePreSetTestInstructionAttributes": [
              {
                "TestInstructionAttributeType": "TEXTBOX",
                "TestInstructionAttributeUuid": "5cc2566e-bf27-44fd-818c-4591fa9c603c",
                "TestInstructionAttributeName": "This attributes holds a Text but attribute is missing",
                "AttributeValueAsString": "LambertJonas",
                "AttributeValueUuid": "7fb2566e-bf27-44fd-818c-4591fa9c603c"
              }
            ]
          }
        ]
      },
      "ImmatureSubTestCaseModel": {
        "FirstImmatureElementUuid": "",
        "TestCaseModelElements": [
          {
            "OriginalElementUuid": "f81b9734-5dce-43c9-8d77-3368940cf126",
            "OriginalElementName": "Emtpy serial processed TestInstructionsContainer",
            "MatureElementUuid": "",
            "PreviousElementUuid": "0883d538-1cff-4be1-ba1f-4dc1f68f6242",
            "NextElementUuid": "0883d538-1cff-4be1-ba1f-4dc1f68f6242",
            "FirstChildElementUuid": "0883d538-1cff-4be1-ba1f-4dc1f68f6242",
            "ParentElementUuid": "f81b9734-5dce-43c9-8d77-3368940cf126",
            "TestCaseModelElementType": "B10_BOND",
            "CurrentElementModelElement": "0883d538-1cff-4be1-ba1f-4dc1f68f6242"
          },
          {
            "OriginalElementUuid": "f81b9734-5dce-43c9-8d77-3368940cf126",
            "OriginalElementName": "Emtpy serial processed TestInstructionsContainer",
            "MatureElementUuid": "",
            "PreviousElementUuid": "f81b9734-5dce-43c9-8d77-3368940cf126",
            "NextElementUuid": "f81b9734-5dce-43c9-8d77-3368940cf126",
            "FirstChildElementUuid": "0883d538-1cff-4be1-ba1f-4dc1f68f6242",
            "ParentElementUuid": "f81b9734-5dce-43c9-8d77-3368940cf126",
            "TestCaseModelElementType": "TIC_TESTINSTRUCTIONCONTAINER",
            "CurrentElementModelElement": "f81b9734-5dce-43c9-8d77-3368940cf126"
          }
        ]
      }
    }
  ],
  "ackNackResponse": {
    "AckNack": true,
    "Comments": "",
    "ErrorCodes": [],
    "ProtoFileVersionUsedByClient": "VERSION_0_3"
  }
}`

// *** Expected result, regarding TestInstructions, when asking for Available Building Blocks

// All TestDomains, TestInstructionTypes and TestInstructions
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_001 string = "map[2f130d7e-f8aa-466f-b29d-0fb63608c1a6:{Just the name [2f130d7] Just the name (Custody Arrangement) [2f130d7] 2f130d7e-f8aa-466f-b29d-0fb63608c1a6 Just the name 1} 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb:{The type of the TestInstruction [513dd8f] The type of the TestInstruction (Can not be pinned) [513dd8f] 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb The type of the TestInstruction 0} 78a97c41-a098-4122-88d2-01ed4b6c4844:{Custody Arrangement [78a97c4] Custody Arrangement (Can not be pinned) [78a97c4] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement 0}]"

// The testCaseModel with structure adjusted to a tree-view, Only TestInstruction-part
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_002 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:map[513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb:map[2f130d7e-f8aa-466f-b29d-0fb63608c1a6:{Just the name [2f130d7] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb The type of the TestInstruction 2f130d7e-f8aa-466f-b29d-0fb63608c1a6 Just the name}]]]"

// *** Expected result, regarding TestInstructionContainerss, when asking for Available Building Blocks

// All TestDomains, TestInstructionContainerTypes and TestInstructionContainers
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_003 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:{Custody Arrangement [78a97c4] Custody Arrangement (Can not be pinned) [78a97c4] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement 0} aa1b9734-5dce-43c9-8d77-3368940cf126:{Emtpy parallelled processed Turbo TestInstructionsContainer [aa1b973] Emtpy parallelled processed Turbo TestInstructionsContainer (Custody Arrangement) [aa1b973] aa1b9734-5dce-43c9-8d77-3368940cf126 Emtpy parallelled processed Turbo TestInstructionsContainer 2} b107bdd9-4152-4020-b3f0-fc750b45885e:{Base containers [b107bdd] Base containers (Can not be pinned) [b107bdd] b107bdd9-4152-4020-b3f0-fc750b45885e Base containers 0} ca07bdd9-4152-4020-b3f0-fc750b45885e:{CA Base containers [ca07bdd] CA Base containers (Can not be pinned) [ca07bdd] ca07bdd9-4152-4020-b3f0-fc750b45885e CA Base containers 0} e107bdd9-4152-4020-b3f0-fc750b45885e:{Emtpy parallelled processed TestInstructionsContainer [e107bdd] Emtpy parallelled processed TestInstructionsContainer (Fenix) [e107bdd] e107bdd9-4152-4020-b3f0-fc750b45885e Emtpy parallelled processed TestInstructionsContainer 2} e81b9734-5dce-43c9-8d77-3368940cf126:{Fenix [e81b973] Fenix (Can not be pinned) [e81b973] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix 0} f81b9734-5dce-43c9-8d77-3368940cf126:{Emtpy serial processed TestInstructionsContainer [f81b973] Emtpy serial processed TestInstructionsContainer (Fenix) [f81b973] f81b9734-5dce-43c9-8d77-3368940cf126 Emtpy serial processed TestInstructionsContainer 2}]"

// The testCaseModel with structure adjusted to a tree-view, Only TestInstruction-part
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_004 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:map[ca07bdd9-4152-4020-b3f0-fc750b45885e:map[aa1b9734-5dce-43c9-8d77-3368940cf126:{Emtpy parallelled processed Turbo TestInstructionsContainer [aa1b973] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement ca07bdd9-4152-4020-b3f0-fc750b45885e CA Base containers aa1b9734-5dce-43c9-8d77-3368940cf126 Emtpy parallelled processed Turbo TestInstructionsContainer}]] e81b9734-5dce-43c9-8d77-3368940cf126:map[b107bdd9-4152-4020-b3f0-fc750b45885e:map[e107bdd9-4152-4020-b3f0-fc750b45885e:{Emtpy parallelled processed TestInstructionsContainer [e107bdd] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix b107bdd9-4152-4020-b3f0-fc750b45885e Base containers e107bdd9-4152-4020-b3f0-fc750b45885e Emtpy parallelled processed TestInstructionsContainer} f81b9734-5dce-43c9-8d77-3368940cf126:{Emtpy serial processed TestInstructionsContainer [f81b973] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix b107bdd9-4152-4020-b3f0-fc750b45885e Base containers f81b9734-5dce-43c9-8d77-3368940cf126 Emtpy serial processed TestInstructionsContainer}]]]"

// *** Expected result, regarding TestInstructions & TestInstructionContainerss, when asking for Available Building Blocks

// All TestDomains, TestInstructionTypes, TestInstructions, TestInstructionContainerTypes and TestInstructionContainers
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_005 string = "map[2f130d7e-f8aa-466f-b29d-0fb63608c1a6:{Just the name [2f130d7] Just the name (Custody Arrangement) [2f130d7] 2f130d7e-f8aa-466f-b29d-0fb63608c1a6 Just the name 1} 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb:{The type of the TestInstruction [513dd8f] The type of the TestInstruction (Can not be pinned) [513dd8f] 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb The type of the TestInstruction 0} 78a97c41-a098-4122-88d2-01ed4b6c4844:{Custody Arrangement [78a97c4] Custody Arrangement (Can not be pinned) [78a97c4] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement 0} aa1b9734-5dce-43c9-8d77-3368940cf126:{Emtpy parallelled processed Turbo TestInstructionsContainer [aa1b973] Emtpy parallelled processed Turbo TestInstructionsContainer (Custody Arrangement) [aa1b973] aa1b9734-5dce-43c9-8d77-3368940cf126 Emtpy parallelled processed Turbo TestInstructionsContainer 2} b107bdd9-4152-4020-b3f0-fc750b45885e:{Base containers [b107bdd] Base containers (Can not be pinned) [b107bdd] b107bdd9-4152-4020-b3f0-fc750b45885e Base containers 0} ca07bdd9-4152-4020-b3f0-fc750b45885e:{CA Base containers [ca07bdd] CA Base containers (Can not be pinned) [ca07bdd] ca07bdd9-4152-4020-b3f0-fc750b45885e CA Base containers 0} e107bdd9-4152-4020-b3f0-fc750b45885e:{Emtpy parallelled processed TestInstructionsContainer [e107bdd] Emtpy parallelled processed TestInstructionsContainer (Fenix) [e107bdd] e107bdd9-4152-4020-b3f0-fc750b45885e Emtpy parallelled processed TestInstructionsContainer 2} e81b9734-5dce-43c9-8d77-3368940cf126:{Fenix [e81b973] Fenix (Can not be pinned) [e81b973] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix 0} f81b9734-5dce-43c9-8d77-3368940cf126:{Emtpy serial processed TestInstructionsContainer [f81b973] Emtpy serial processed TestInstructionsContainer (Fenix) [f81b973] f81b9734-5dce-43c9-8d77-3368940cf126 Emtpy serial processed TestInstructionsContainer 2}]"

// The testCaseModel with structure adjusted to a tree-view, Only TestInstruction-part
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_006 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:map[513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb:map[2f130d7e-f8aa-466f-b29d-0fb63608c1a6:{Just the name [2f130d7] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb The type of the TestInstruction 2f130d7e-f8aa-466f-b29d-0fb63608c1a6 Just the name}]]]"

// The testCaseModel with structure adjusted to a tree-view, Only TestInstruction-part
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_007 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:map[ca07bdd9-4152-4020-b3f0-fc750b45885e:map[aa1b9734-5dce-43c9-8d77-3368940cf126:{Emtpy parallelled processed Turbo TestInstructionsContainer [aa1b973] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement ca07bdd9-4152-4020-b3f0-fc750b45885e CA Base containers aa1b9734-5dce-43c9-8d77-3368940cf126 Emtpy parallelled processed Turbo TestInstructionsContainer}]] e81b9734-5dce-43c9-8d77-3368940cf126:map[b107bdd9-4152-4020-b3f0-fc750b45885e:map[e107bdd9-4152-4020-b3f0-fc750b45885e:{Emtpy parallelled processed TestInstructionsContainer [e107bdd] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix b107bdd9-4152-4020-b3f0-fc750b45885e Base containers e107bdd9-4152-4020-b3f0-fc750b45885e Emtpy parallelled processed TestInstructionsContainer} f81b9734-5dce-43c9-8d77-3368940cf126:{Emtpy serial processed TestInstructionsContainer [f81b973] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix b107bdd9-4152-4020-b3f0-fc750b45885e Base containers f81b9734-5dce-43c9-8d77-3368940cf126 Emtpy serial processed TestInstructionsContainer}]]]"

// ******* END ABB001 *******

// Pinned Building Blocks; PBB

// ******* START PBB001 *******
// gRPC-result when asking for Pinned Building Blocks
var PinnedTestInstructionsAndTestInstructionsContainersRespons_PBB001 string = `{
  "AvailablePinnedTestInstructions": [
    {
      "TestInstructionUuid": "2f130d7e-f8aa-466f-b29d-0fb63608c1a6",
      "TestInstructionName": "TestInstructionName 1"
    }
  ],
  "AvailablePinnedPreCreatedTestInstructionContainers": [
    {
      "TestInstructionContainerUuid": "f81b9734-5dce-43c9-8d77-3368940cf126",
      "TestInstructionContainerName": "TestInstructionContainerName 1"
    },
    {
      "TestInstructionContainerUuid": "e107bdd9-4152-4020-b3f0-fc750b45885e",
      "TestInstructionContainerName": "TestInstructionContainerName"
    }
  ],
  "ackNackResponse": {
    "AckNack": true,
    "Comments": "",
    "ErrorCodes": [],
    "ProtoFileVersionUsedByClient": "VERSION_0_3"
  }
}`

// All Pinned TestInstructions
var TestInstructionsAndTestInstructionsContainersRespons_PBB001_ExpectedResultInModel_001 string = "map[Just the name (Custody Arrangement) [2f130d7]:{2f130d7e-f8aa-466f-b29d-0fb63608c1a6 1}]"

// All Pinned TestInstructionContainers
var TestInstructionsAndTestInstructionsContainersRespons_PBB001_ExpectedResultInModel_002 string = "map[Emtpy parallelled processed TestInstructionsContainer (Fenix) [e107bdd]:{e107bdd9-4152-4020-b3f0-fc750b45885e 2} Emtpy serial processed TestInstructionsContainer (Fenix) [f81b973]:{f81b9734-5dce-43c9-8d77-3368940cf126 2}]"

// All Pinned TestInstructions and TestInstructionContainers
var TestInstructionsAndTestInstructionsContainersRespons_PBB001_ExpectedResultInModel_003 string = "map[Emtpy parallelled processed TestInstructionsContainer (Fenix) [e107bdd]:{e107bdd9-4152-4020-b3f0-fc750b45885e 2} Emtpy serial processed TestInstructionsContainer (Fenix) [f81b973]:{f81b9734-5dce-43c9-8d77-3368940cf126 2} Just the name (Custody Arrangement) [2f130d7]:{2f130d7e-f8aa-466f-b29d-0fb63608c1a6 1}]"

// ******* END PBB001 *******

// ******* START PBB002 *******
// Error text after trying to pin an non existing Building Block
var TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_001 string = "NonExistingBuildingBlock is missing among nodes i map"

// Error message after trying to pin an already pinned Building Block
var TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_002 string = "building block is already pinned"

// Error message after Validate that a non pinned Building Block can be pinned
var TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_003 string = "<nil>"

// Map after pinning a Building Block that can be pinned
var TestInstructionsAndTestInstructionsContainersRespons_PBB002_ExpectedResultInModel_004 string = "map[Emtpy parallelled processed TestInstructionsContainer (Fenix) [e107bdd]:{e107bdd9-4152-4020-b3f0-fc750b45885e 2} Emtpy parallelled processed Turbo TestInstructionsContainer (Custody Arrangement) [aa1b973]:{aa1b9734-5dce-43c9-8d77-3368940cf126 2} Emtpy serial processed TestInstructionsContainer (Fenix) [f81b973]:{f81b9734-5dce-43c9-8d77-3368940cf126 2} Just the name (Custody Arrangement) [2f130d7]:{2f130d7e-f8aa-466f-b29d-0fb63608c1a6 1}]"

// ******* END PBB002 *******

// ******* START PBB003 *******
// Error text after trying to unpin an non existing Building Block
var TestInstructionsAndTestInstructionsContainersRespons_PBB003_ExpectedResultInModel_001 string = "NonExistingBuildingBlock is missing among nodes i map"

// Error message after trying toun pin an Building Block that doesn't exist among pinned building blocks
var TestInstructionsAndTestInstructionsContainersRespons_PBB003_ExpectedResultInModel_002 string = "building block is already pinned"

// Error message after Validate that a pinned Building Block can be unpinned
var TestInstructionsAndTestInstructionsContainersRespons_PBB003_ExpectedResultInModel_003 string = "<nil>"

// Map after unpinning a Building Block that was pinned
var TestInstructionsAndTestInstructionsContainersRespons_PBB003_ExpectedResultInModel_004 string = "map[Emtpy parallelled processed TestInstructionsContainer (Fenix) [e107bdd]:{e107bdd9-4152-4020-b3f0-fc750b45885e 2} Emtpy parallelled processed Turbo TestInstructionsContainer (Custody Arrangement) [aa1b973]:{aa1b9734-5dce-43c9-8d77-3368940cf126 2} Emtpy serial processed TestInstructionsContainer (Fenix) [f81b973]:{f81b9734-5dce-43c9-8d77-3368940cf126 2} Just the name (Custody Arrangement) [2f130d7]:{2f130d7e-f8aa-466f-b29d-0fb63608c1a6 1}]"

// ******* END PBB003 *******

const loggingLevelForDebug = logrus.ErrorLevel // InfoLevel

// Init the logger for UnitTests
func InitLoggerForTest(filename string) (myTestLogger *logrus.Logger) {
	myTestLogger = logrus.StandardLogger()

	//log.Println("'loggingLevelForDebug': ", loggingLevelForDebug)

	logrus.SetLevel(loggingLevelForDebug)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  true,
	})

	//If no file then set standard out

	if filename == "" {
		myTestLogger.Out = os.Stdout

	} else {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			myTestLogger.Out = file
		} else {
			log.Println("Failed to log to file, using default stderr")
		}
	}

	// Should only be done from init functions
	//grpclog.SetLoggerV2(grpclog.NewLoggerV2(logger.Out, logger.Out, logger.Out))

	return myTestLogger
}

// ********************************************************************************************************************
// Check if testdata is using correct proto-file version
func IsTestDataUsingCorrectTestDataProtoFileVersion(usedProtoFileVersion fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum) (returnMessage *fenixTestCaseBuilderServerGrpcApi.AckNackResponse) {

	var clientUseCorrectProtoFileVersion bool
	var protoFileExpected fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum
	var protoFileUsed fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum

	protoFileUsed = usedProtoFileVersion
	protoFileExpected = fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum(getHighestFenixTestDataProtoFileVersion())

	// Check if correct proto files is used
	if protoFileExpected == protoFileUsed {
		clientUseCorrectProtoFileVersion = true
	} else {
		clientUseCorrectProtoFileVersion = false
	}

	// Check if Client is using correct proto files version
	if clientUseCorrectProtoFileVersion == false {
		// Not correct proto-file version is used

		// Set Error codes to return message
		var errorCodes []fenixTestCaseBuilderServerGrpcApi.ErrorCodesEnum
		var errorCode fenixTestCaseBuilderServerGrpcApi.ErrorCodesEnum

		errorCode = fenixTestCaseBuilderServerGrpcApi.ErrorCodesEnum_ERROR_WRONG_PROTO_FILE_VERSION
		errorCodes = append(errorCodes, errorCode)

		// Create Return message
		returnMessage = &fenixTestCaseBuilderServerGrpcApi.AckNackResponse{
			AckNack:    false,
			Comments:   "Wrong proto file used. Expected: '" + protoFileExpected.String() + "', but got: '" + protoFileUsed.String() + "'",
			ErrorCodes: errorCodes,
		}

		return returnMessage

	} else {
		return nil
	}

}

// ********************************************************************************************************************
var highestFenixProtoFileVersion int32

// Get the highest FenixProtoFileVersionEnumeration
func getHighestFenixTestDataProtoFileVersion() int32 {

	// Check if there already is a 'highestFenixProtoFileVersion' saved, if so use that one
	if highestFenixProtoFileVersion > 0 {
		return highestFenixProtoFileVersion
	}

	// Find the highest value for proto-file version
	var maxValue int32
	maxValue = 0

	for _, v := range fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	highestFenixProtoFileVersion = maxValue

	return highestFenixProtoFileVersion
}
