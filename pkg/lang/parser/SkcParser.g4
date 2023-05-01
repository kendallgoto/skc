parser grammar SkcParser;
options { tokenVocab=SkcLexer; }

enter
    :   line*
    |   EOF
    ;
line
    :   statement ((LineTerminal WS?)|EOF)
    ;
statement
    :   sayStatement
    |   conditionalStatement
    |   whileStatement
    ;
conditionalStatement
    :   If WS condition COMMA WS Then WS statement
    ;

whileStatement
    :   While WS condition COMMA WS statement
    ;

condition
    :   literal WS Is (WS (Not WS)? equality)? WS literal
    ;
equality
    :   Equal
    |   GreaterThan
    |   LessThan
    |   GreaterThan
    |   GreaterThanOrEqual
    |   LessThanOrEqual
    ;
sayStatement
    :   Say WS literal (WS OutputDestination WS (OutputStdout | (OutputFile WS StringLiteral)))?
    ;
literal
    :   StringLiteral
    |   Constant
    ;
outputType
    :   OutputStdout
    |   OutputFile WS StringLiteral
    ;
