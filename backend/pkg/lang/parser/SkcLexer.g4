lexer grammar SkcLexer;

SayCommand : S A Y;

LineTerminal : '.';
WS : [ \t\n\r\f]+;
StringSeperator : '"';
OutputDestination: T O;
OutputStdout: M E;
OutputFile: F I L E;

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
// https://github.com/tunnelvisionlabs/antlr4/blob/master/doc/case-insensitive-lexing.md
fragment A : [aA]; // match either an 'a' or 'A'
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];

ErrorCharacter : . ;
