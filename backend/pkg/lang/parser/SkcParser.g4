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
    ;
conditionalStatement
    :   If WS condition COMMA WS Then WS statement
    ;
condition
    :   literal WS Is (WS equality)? WS literal
    ;
equality
    :   Equal
    |   GreaterThan
    |   LessThan
    ;
sayStatement
    :   Say WS literal (WS outputTo)?
    ;
outputTo
    :   OutputDestination WS outputType
    ;
literal
    :   StringLiteral
    |   Constant
    ;
outputType
    :   OutputStdout
    |   OutputFile WS StringLiteral
    ;
