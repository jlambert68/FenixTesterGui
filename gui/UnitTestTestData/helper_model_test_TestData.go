package UnitTestTestData

// Availalbe Buoilding Blocks; ABB

// ******* START ABB001 *******
// Result when asking for Available Building Blocks
var TestInstructionsAndTestInstructionsContainersRespons_ABB001 string = `{
		"TestInstructionMessages": [
	{
	"DomainUuid": "78a97c41-a098-4122-88d2-01ed4b6c4844",
	"DomainName": "Custody Arrangement",
	"TestInstructionUuid": "2f130d7e-f8aa-466f-b29d-0fb63608c1a6",
	"TestInstructionName": "Just the name",
	"TestInstructionTypeUuid": "513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb",
	"TestInstructionTypeName": "The type of the TestInstruction",
	"TestInstructionDescription": "En vanlig typ",
	"TestInstructionMouseOverText": "This will be shown when hovering above this TestInstruction",
	"Deprecated": false,
	"Enabled": true,
	"MajorVersionNumber": 0,
	"MinorVersionNumber": 1,
	"UpdatedTimeStamp": "2022-04-29T15:42:15Z"
	}
	],
	"TestInstructionContainerMessages": [
	{
	"DomainUuid": "e81b9734-5dce-43c9-8d77-3368940cf126",
	"DomainName": "Fenix",
	"TestInstructionContainerUuid": "f81b9734-5dce-43c9-8d77-3368940cf126",
	"TestInstructionContainerName": "Emtpy serial processed TestInstructionsContainer",
	"TestInstructionContainerTypeUuid": "b107bdd9-4152-4020-b3f0-fc750b45885e",
	"TestInstructionContainerTypeName": "Base containers",
	"TestInstructionContainerDescription": "Children of this container is processed in serial",
	"TestInstructionContainerMouseOverText": "Children of this container is processed in serial",
	"Deprecated": false,
	"Enabled": true,
	"MajorVersionNumber": 0,
	"MinorVersionNumber": 1,
	"UpdatedTimeStamp": "2022-05-02T10:10:07Z",
	"TestInstructionContainerExecutionType": "SERIAL_PROCESSED",
	"TestInstructionContainerChildren": []
	},
	{
	"DomainUuid": "e81b9734-5dce-43c9-8d77-3368940cf126",
	"DomainName": "Fenix",
	"TestInstructionContainerUuid": "e107bdd9-4152-4020-b3f0-fc750b45885e",
	"TestInstructionContainerName": "Emtpy parallelled processed TestInstructionsContainer",
	"TestInstructionContainerTypeUuid": "b107bdd9-4152-4020-b3f0-fc750b45885e",
	"TestInstructionContainerTypeName": "Base containers",
	"TestInstructionContainerDescription": "Children of this container is processed in parallel",
	"TestInstructionContainerMouseOverText": "Children of this container is processed in parallel",
	"Deprecated": false,
	"Enabled": true,
	"MajorVersionNumber": 0,
	"MinorVersionNumber": 1,
	"UpdatedTimeStamp": "2022-05-02T10:08:28Z",
	"TestInstructionContainerExecutionType": "PARALLELLED_PROCESSED",
	"TestInstructionContainerChildren": []
	},
	{
	"DomainUuid": "78a97c41-a098-4122-88d2-01ed4b6c4844",
	"DomainName": "Custody Arrangement",
	"TestInstructionContainerUuid": "aa1b9734-5dce-43c9-8d77-3368940cf126",
	"TestInstructionContainerName": "Emtpy parallelled processed Turbo TestInstructionsContainer",
	"TestInstructionContainerTypeUuid": "ca07bdd9-4152-4020-b3f0-fc750b45885e",
	"TestInstructionContainerTypeName": "CA Base containers",
	"TestInstructionContainerDescription": "Children of this CA-container is processed in parallel",
	"TestInstructionContainerMouseOverText": "Children of this CA-container is processed in parallel",
	"Deprecated": false,
	"Enabled": true,
	"MajorVersionNumber": 0,
	"MinorVersionNumber": 1,
	"UpdatedTimeStamp": "2022-06-16T16:09:43Z",
	"TestInstructionContainerExecutionType": "PARALLELLED_PROCESSED",
	"TestInstructionContainerChildren": []
	}
	],
	"ackNackResponse": {
	"AckNack": true,
	"Comments": "",
	"ErrorCodes": []
	}
	}`

// *** Expected result, regarding TestInstructions, when asking for Available Building Blocks

// All TestDomains, TestInstructionTypes and TestInstructions
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_001 string = "map[2f130d7e-f8aa-466f-b29d-0fb63608c1a6:{Just the name [2f130d7] 2f130d7e-f8aa-466f-b29d-0fb63608c1a6 Just the name} 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb:{The type of the TestInstruction [513dd8f] 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb The type of the TestInstruction} 78a97c41-a098-4122-88d2-01ed4b6c4844:{Custody Arrangement [78a97c4] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement}]"

// The model with structure adjusted to a tree-view, Only TestInstruction-part
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_002 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:map[513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb:map[2f130d7e-f8aa-466f-b29d-0fb63608c1a6:{Just the name [2f130d7] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb The type of the TestInstruction 2f130d7e-f8aa-466f-b29d-0fb63608c1a6 Just the name}]]]"

// *** Expected result, regarding TestInstructionContainerss, when asking for Available Building Blocks

// All TestDomains, TestInstructionContainerTypes and TestInstructionContainers
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_003 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:{Custody Arrangement [78a97c4] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement} aa1b9734-5dce-43c9-8d77-3368940cf126:{Emtpy parallelled processed Turbo TestInstructionsContainer [aa1b973] aa1b9734-5dce-43c9-8d77-3368940cf126 Emtpy parallelled processed Turbo TestInstructionsContainer} b107bdd9-4152-4020-b3f0-fc750b45885e:{Base containers [b107bdd] b107bdd9-4152-4020-b3f0-fc750b45885e Base containers} ca07bdd9-4152-4020-b3f0-fc750b45885e:{CA Base containers [ca07bdd] ca07bdd9-4152-4020-b3f0-fc750b45885e CA Base containers} e107bdd9-4152-4020-b3f0-fc750b45885e:{Emtpy parallelled processed TestInstructionsContainer [e107bdd] e107bdd9-4152-4020-b3f0-fc750b45885e Emtpy parallelled processed TestInstructionsContainer} e81b9734-5dce-43c9-8d77-3368940cf126:{Fenix [e81b973] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix} f81b9734-5dce-43c9-8d77-3368940cf126:{Emtpy serial processed TestInstructionsContainer [f81b973] f81b9734-5dce-43c9-8d77-3368940cf126 Emtpy serial processed TestInstructionsContainer}]"

// The model with structure adjusted to a tree-view, Only TestInstruction-part
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_004 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:map[ca07bdd9-4152-4020-b3f0-fc750b45885e:map[aa1b9734-5dce-43c9-8d77-3368940cf126:{Emtpy parallelled processed Turbo TestInstructionsContainer [aa1b973] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement ca07bdd9-4152-4020-b3f0-fc750b45885e CA Base containers aa1b9734-5dce-43c9-8d77-3368940cf126 Emtpy parallelled processed Turbo TestInstructionsContainer}]] e81b9734-5dce-43c9-8d77-3368940cf126:map[b107bdd9-4152-4020-b3f0-fc750b45885e:map[e107bdd9-4152-4020-b3f0-fc750b45885e:{Emtpy parallelled processed TestInstructionsContainer [e107bdd] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix b107bdd9-4152-4020-b3f0-fc750b45885e Base containers e107bdd9-4152-4020-b3f0-fc750b45885e Emtpy parallelled processed TestInstructionsContainer} f81b9734-5dce-43c9-8d77-3368940cf126:{Emtpy serial processed TestInstructionsContainer [f81b973] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix b107bdd9-4152-4020-b3f0-fc750b45885e Base containers f81b9734-5dce-43c9-8d77-3368940cf126 Emtpy serial processed TestInstructionsContainer}]]]"

// *** Expected result, regarding TestInstructions & TestInstructionContainerss, when asking for Available Building Blocks

// All TestDomains, TestInstructionTypes, TestInstructions, TestInstructionContainerTypes and TestInstructionContainers
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_005 string = "map[2f130d7e-f8aa-466f-b29d-0fb63608c1a6:{Just the name [2f130d7] 2f130d7e-f8aa-466f-b29d-0fb63608c1a6 Just the name} 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb:{The type of the TestInstruction [513dd8f] 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb The type of the TestInstruction} 78a97c41-a098-4122-88d2-01ed4b6c4844:{Custody Arrangement [78a97c4] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement} aa1b9734-5dce-43c9-8d77-3368940cf126:{Emtpy parallelled processed Turbo TestInstructionsContainer [aa1b973] aa1b9734-5dce-43c9-8d77-3368940cf126 Emtpy parallelled processed Turbo TestInstructionsContainer} b107bdd9-4152-4020-b3f0-fc750b45885e:{Base containers [b107bdd] b107bdd9-4152-4020-b3f0-fc750b45885e Base containers} ca07bdd9-4152-4020-b3f0-fc750b45885e:{CA Base containers [ca07bdd] ca07bdd9-4152-4020-b3f0-fc750b45885e CA Base containers} e107bdd9-4152-4020-b3f0-fc750b45885e:{Emtpy parallelled processed TestInstructionsContainer [e107bdd] e107bdd9-4152-4020-b3f0-fc750b45885e Emtpy parallelled processed TestInstructionsContainer} e81b9734-5dce-43c9-8d77-3368940cf126:{Fenix [e81b973] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix} f81b9734-5dce-43c9-8d77-3368940cf126:{Emtpy serial processed TestInstructionsContainer [f81b973] f81b9734-5dce-43c9-8d77-3368940cf126 Emtpy serial processed TestInstructionsContainer}]"

// The model with structure adjusted to a tree-view, Only TestInstruction-part
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_006 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:map[513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb:map[2f130d7e-f8aa-466f-b29d-0fb63608c1a6:{Just the name [2f130d7] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement 513dd8fb-a0bb-4738-9a0b-b7eaf7bb8adb The type of the TestInstruction 2f130d7e-f8aa-466f-b29d-0fb63608c1a6 Just the name}]]]"

// The model with structure adjusted to a tree-view, Only TestInstruction-part
var TestInstructionsAndTestInstructionsContainersRespons_ABB001_ExpectedResultInModel_007 string = "map[78a97c41-a098-4122-88d2-01ed4b6c4844:map[ca07bdd9-4152-4020-b3f0-fc750b45885e:map[aa1b9734-5dce-43c9-8d77-3368940cf126:{Emtpy parallelled processed Turbo TestInstructionsContainer [aa1b973] 78a97c41-a098-4122-88d2-01ed4b6c4844 Custody Arrangement ca07bdd9-4152-4020-b3f0-fc750b45885e CA Base containers aa1b9734-5dce-43c9-8d77-3368940cf126 Emtpy parallelled processed Turbo TestInstructionsContainer}]] e81b9734-5dce-43c9-8d77-3368940cf126:map[b107bdd9-4152-4020-b3f0-fc750b45885e:map[e107bdd9-4152-4020-b3f0-fc750b45885e:{Emtpy parallelled processed TestInstructionsContainer [e107bdd] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix b107bdd9-4152-4020-b3f0-fc750b45885e Base containers e107bdd9-4152-4020-b3f0-fc750b45885e Emtpy parallelled processed TestInstructionsContainer} f81b9734-5dce-43c9-8d77-3368940cf126:{Emtpy serial processed TestInstructionsContainer [f81b973] e81b9734-5dce-43c9-8d77-3368940cf126 Fenix b107bdd9-4152-4020-b3f0-fc750b45885e Base containers f81b9734-5dce-43c9-8d77-3368940cf126 Emtpy serial processed TestInstructionsContainer}]]]"

// ******* END ABB001 *******
