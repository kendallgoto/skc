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
    ;
sayStatement
    :   SayCommand WS literal (WS outputTo)?
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
