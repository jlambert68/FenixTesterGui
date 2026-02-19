# common_merkleTreeFunctions.go

## File Overview
- Path: `common_code/common_merkleTreeFunctions.go`
- Package: `sharedCode`
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
### HashValues
- Signature: `func HashValues(valuesToHash []string, isNonHashValue bool) string`
- Exported: `true`
- Control-flow features: `if, for/range`
- Selector calls: `strconv.Itoa`, `sort.Strings`, `sha256.New`, `hash.Write`, `hex.EncodeToString`, `hash.Sum`

### setFromList
- Signature: `func setFromList(list []string) set []string`
- Exported: `false`
- Control-flow features: `if, for/range`

### uniqueGotaSeries
- Signature: `func uniqueGotaSeries(s series.Series) series.Series`
- Exported: `false`
- Control-flow features: `none detected`
- Internal calls: `setFromList`
- Selector calls: `series.New`, `s.Records`, `s.Type`

### uniqueGotaSeriesAsStringArray
- Signature: `func uniqueGotaSeriesAsStringArray(s series.Series) []string`
- Exported: `false`
- Control-flow features: `none detected`
- Internal calls: `uniqueGotaSeries`

### hashChildrenAndWriteToDataStore
- Signature: `func hashChildrenAndWriteToDataStore(level int, currentMerklePath string, valuesToHash []string, isEndLeafNode bool, currentMerkleFilterPath string) string`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `HashSingleValue`
- Selector calls: `sort.Strings`, `sha256.New`, `hash.Write`, `hex.EncodeToString`, `hash.Sum`, `dataframe.New`, `series.New`, `merkleTreeDataFrame.RBind`

### recursiveTreeCreator
- Signature: `func recursiveTreeCreator(level int, currentMerkleFilterPath string, dataFrameToWorkOn dataframe.DataFrame, currentMerklePath string) (string, string)`
- Exported: `false`
- Control-flow features: `if, for/range`
- Internal calls: `uniqueGotaSeriesAsStringArray`, `hashChildrenAndWriteToDataStore`, `recursiveTreeCreator`
- Selector calls: `strings.Index`, `dataFrameToWorkOn.Arrange`, `dataframe.Sort`, `dataFrameToWorkOn.Col`, `dataFrameToWorkOn.Filter`, `log.Fatalln`, `strings.HasSuffix`, `strings.LastIndex`

### LoadAndProcessFile
- Signature: `func LoadAndProcessFile(fileToprocess string) (string, dataframe.DataFrame, dataframe.DataFrame)`
- Exported: `true`
- Control-flow features: `if, for/range, defer`
- Doc: Process incoming csv file and create MerkleRootHash and MerkleTree
- Internal calls: `HashValues`, `recursiveTreeCreator`
- Selector calls: `os.Open`, `log.Fatal`, `irisCsv.Close`, `dataframe.ReadCSV`, `dataframe.WithDelimiter`, `dataframe.HasHeader`, `df.Arrange`, `dataframe.Sort`

### calculateMerkleHashFromMerkleTreeLeafNodes
- Signature: `func calculateMerkleHashFromMerkleTreeLeafNodes(merkleLevel int, merkleTreeLeafNodes dataframe.DataFrame, maxMerkleLevel int) merkleHash string`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Calculate MerkleHash from leaf nodes in MerkleTree
- Internal calls: `uniqueGotaSeriesAsStringArray`, `calculateMerkleHashFromMerkleTreeLeafNodes`, `HashValues`
- Selector calls: `merkleTreeLeafNodes.Col`, `merkleTreeLeafNodes.Arrange`, `dataframe.Sort`, `merkleTreeLeafNodes.Nrow`, `merkleTreeLeafNodes.Ncol`, `merkleTreeLeafNodes.Elem`, `strings.Index`, `merkleTreeLeafNodes.Filter`

### CalculateMerkleHashFromMerkleTree
- Signature: `func CalculateMerkleHashFromMerkleTree(merkleTree dataframe.DataFrame) merkleHash string`
- Exported: `true`
- Control-flow features: `if`
- Doc: CalculateMerkleHashFromMerkleTree Calculate MerkleHash from leaf nodes in MerkleTree
- Internal calls: `int`, `calculateMerkleHashFromMerkleTreeLeafNodes`
- Selector calls: `merkleTree.Col`, `merkleTree.Filter`, `merkleTreeLeafNodes.Nrow`, `merkleTreeLeafNodes.Mutate`, `series.New`

### ExtractMerkleRootHashFromMerkleTree
- Signature: `func ExtractMerkleRootHashFromMerkleTree(merkleTree dataframe.DataFrame) merkleRootHash string`
- Exported: `true`
- Control-flow features: `if`
- Doc: ExtractMerkleRootHashFromMerkleTree Retrieve MerkleRootHashFromMerkleTree
- Internal calls: `int`, `uniqueGotaSeriesAsStringArray`
- Selector calls: `merkleTree.Col`, `merkleTree.Filter`, `merkleTreeRoot.Col`

### ExtractValuesFromFilterPath
- Signature: `func ExtractValuesFromFilterPath(merklFilterPath string) merkleFilterColumns []string`
- Exported: `true`
- Control-flow features: `if, for/range`
- Doc: Extract all MerkleFIlterCOlumns from a MerkleFilterPath
- Selector calls: `strings.Index`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
