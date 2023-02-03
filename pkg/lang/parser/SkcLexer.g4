lexer grammar SkcLexer;
options { caseInsensitive = true; }

Say : 'say';

LineTerminal : '.';
WS : [ \t\n\r\f]+;
StringSeperator : '"';
OutputDestination: 'to';
OutputStdout: 'me';
OutputFile: 'file';
If: 'if';
COMMA : ',';
Then :'then';
Is : 'is';
Equal : 'equal' WS 'to';
Not : 'not';
GreaterThan : 'greater' WS 'than';
LessThan : 'less' WS 'than';
GreaterThanOrEqual : 'greater' WS 'than' WS 'or' WS 'equal' WS 'to';
LessThanOrEqual : 'less' WS 'than' WS 'or' WS 'equal' WS 'to';
StringLiteral
    :   StringSeperator CharSequence? StringSeperator
    ;

fragment
CharSequence
    :  Char+
    ;
fragment
Char
    :   ~["\\\r\n]
    ;

Constant
    :   Num+
    ;

fragment
Num
    :   [0-9]
    ;

ErrorCharacter : . ;
