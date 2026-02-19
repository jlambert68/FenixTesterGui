# common_merkleTreeFunctions.go

## File Overview
- Path: `common_code/common_merkleTreeFunctions.go`
- Package: `sharedCode`
- Generated: `2026-02-19T14:23:17+01:00`
- Functions/Methods: `11`
- Imports: `9`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CalculateMerkleHashFromMerkleTree`
- `ExtractMerkleRootHashFromMerkleTree`
- `ExtractValuesFromFilterPath`
- `HashValues`
- `LoadAndProcessFile`

## Imports
- `crypto/sha256`
- `encoding/hex`
- `github.com/go-gota/gota/dataframe`
- `github.com/go-gota/gota/series`
- `log`
- `os`
- `sort`
- `strconv`
- `strings`

## Declared Types
- None

## Declared Constants
- None

## Declared Variables
- `changedFilesMerkleTreeDataFrame`
- `merkleTreeDataFrame`

## Functions and Methods
### CalculateMerkleHashFromMerkleTree
- Signature: `func CalculateMerkleHashFromMerkleTree(merkleTree dataframe.DataFrame) merkleHash string`
- Exported: `true`
- Control-flow features: `if`
- Doc: CalculateMerkleHashFromMerkleTree Calculate MerkleHash from leaf nodes in MerkleTree
- Internal calls: `calculateMerkleHashFromMerkleTreeLeafNodes`, `int`
- External calls: `merkleTree.Col`, `merkleTree.Filter`, `merkleTreeLeafNodes.Mutate`, `merkleTreeLeafNodes.Nrow`, `series.New`

### ExtractMerkleRootHashFromMerkleTree
- Signature: `func ExtractMerkleRootHashFromMerkleTree(merkleTree dataframe.DataFrame) merkleRootHash string`
- Exported: `true`
- Control-flow features: `if`
- Doc: ExtractMerkleRootHashFromMerkleTree Retrieve MerkleRootHashFromMerkleTree
- Internal calls: `int`, `uniqueGotaSeriesAsStringArray`
- External calls: `merkleTree.Col`, `merkleTree.Filter`, `merkleTreeRoot.Col`

### ExtractValuesFromFilterPath
- Signature: `func ExtractValuesFromFilterPath(merklFilterPath string) merkleFilterColumns []string`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: Extract all MerkleFIlterCOlumns from a MerkleFilterPath
- External calls: `strings.Index`

### HashValues
- Signature: `func HashValues(valuesToHash []string, isNonHashValue bool) string`
- Exported: `true`
- Control-flow features: `if, for/range`
- External calls: `hash.Sum`, `hash.Write`, `hex.EncodeToString`, `sha256.New`, `sort.Strings`, `strconv.Itoa`

### LoadAndProcessFile
- Signature: `func LoadAndProcessFile(fileToprocess string) (string, dataframe.DataFrame, dataframe.DataFrame)`
- Exported: `true`
- Control-flow features: `if, for/range, defer`
- Doc: Process incoming csv file and create MerkleRootHash and MerkleTree
- Internal calls: `HashValues`, `recursiveTreeCreator`
- External calls: `dataframe.HasHeader`, `dataframe.New`, `dataframe.ReadCSV`, `dataframe.Sort`, `dataframe.WithDelimiter`, `df.Arrange`, `df.Elem`, `df.Mutate`

### calculateMerkleHashFromMerkleTreeLeafNodes
- Signature: `func calculateMerkleHashFromMerkleTreeLeafNodes(merkleLevel int, merkleTreeLeafNodes dataframe.DataFrame, maxMerkleLevel int) merkleHash string`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Calculate MerkleHash from leaf nodes in MerkleTree
- Internal calls: `HashValues`, `calculateMerkleHashFromMerkleTreeLeafNodes`, `uniqueGotaSeriesAsStringArray`
- External calls: `dataframe.Sort`, `log.Fatalln`, `merkleTreeLeafNodes.Arrange`, `merkleTreeLeafNodes.Col`, `merkleTreeLeafNodes.Elem`, `merkleTreeLeafNodes.Filter`, `merkleTreeLeafNodes.Ncol`, `merkleTreeLeafNodes.Nrow`

### hashChildrenAndWriteToDataStore
- Signature: `func hashChildrenAndWriteToDataStore(level int, currentMerklePath string, valuesToHash []string, isEndLeafNode bool, currentMerkleFilterPath string) string`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `HashSingleValue`
- External calls: `dataframe.New`, `hash.Sum`, `hash.Write`, `hex.EncodeToString`, `merkleTreeDataFrame.RBind`, `series.New`, `sha256.New`, `sort.Strings`

### recursiveTreeCreator
- Signature: `func recursiveTreeCreator(level int, currentMerkleFilterPath string, dataFrameToWorkOn dataframe.DataFrame, currentMerklePath string) (string, string)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `hashChildrenAndWriteToDataStore`, `recursiveTreeCreator`, `uniqueGotaSeriesAsStringArray`
- External calls: `dataFrameToWorkOn.Arrange`, `dataFrameToWorkOn.Col`, `dataFrameToWorkOn.Filter`, `dataframe.Sort`, `log.Fatalln`, `strings.HasSuffix`, `strings.Index`, `strings.LastIndex`

### setFromList
- Signature: `func setFromList(list []string) set []string`
- Exported: `false`
- Control-flow features: `if, for/range`

### uniqueGotaSeries
- Signature: `func uniqueGotaSeries(s series.Series) series.Series`
- Exported: `false`
- Control-flow features: `none detected`
- Internal calls: `setFromList`
- External calls: `s.Records`, `s.Type`, `series.New`

### uniqueGotaSeriesAsStringArray
- Signature: `func uniqueGotaSeriesAsStringArray(s series.Series) []string`
- Exported: `false`
- Control-flow features: `none detected`
- Internal calls: `uniqueGotaSeries`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
