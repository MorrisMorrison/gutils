# gutils

## gerror
FailOnError(e error)

## gio
- ReadLinesFromFile(file *os.File, ignoreWhitespace bool) ([]string, error) 
- ReadLinesFromPath(path string, ignoreWhitespace bool) ([]string, error)

## glogger
NewLogger() *Logger
- (l *Logger) Log(level LogLevel, message string)
- (l *Logger) Logf(level LogLevel, format string, v ...interface{})
- (l *Logger) Info(message string)
- (l *Logger) Infof(format string, v ...interface{})
- (l *Logger) Debug(message string)
- (l *Logger) Debugf(format string, v ...interface{})
- (l *Logger) Warning(message string)
- (l *Logger) Warningf(format string, v ...interface{})
- (l *Logger) Error(err error, message string)
- (l *Logger) Errorf(err error, format string, v ...interface{})

## gollections
- Any[T comparable](slice []T, predicate func(T) bool) bool 
- All[T comparable](slice []T, predicate func(T) bool) bool 
- Equals[T comparable](slice1 []T, slice2 []T, comparator func(T, T) bool) bool
- GetColumn(matrix [][]int, colNum int) []int
- GetRow(matrix [][]int, rowNum int) []int
- GetSum(matrix [][]int) int
- ReplaceByMatchingValue(matrix [][]int, comparisonValue int) [][]int
- RemoveRow(slice [][]int, s int) [][]int