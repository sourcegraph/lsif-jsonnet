// language spec: https://jsonnet.org/ref/spec.html
// language tutorial: https://jsonnet.org/learning/tutorial.html
// grammar from: https://gist.github.com/ironchefpython/84380aa60871853dc86719dd598c35e4
// slightly changed the slice rule
grammar Jsonnet;

jsonnet
    : expr EOF
    ;

expr
    : value=(NULL | TRUE | FALSE | SELF | DOLLAR | STRING | NUMBER ) # Value
    | '(' expr ')'                                                   # Parens
    | '{' objinside? '}'                                             # Object
    | '[' ( elems+=expr (',' elems+=expr)* )? ','? ']'               # Array
    | '[' expr ','? forspec+ ']'                                     # ArrayComp
    | expr '.' ID                                                    # Index
    | expr '[' expr ']'                                              # IndexExpr
    | expr '[' startIdx=expr? ':' endIdx=expr? (':' step=expr? )? ']'  # Slice
    | SUPER . ID                                                     # IndexSuper
    | SUPER '[' expr ']'                                             # IndexSuperExpr
    | expr '(' args? ')' TAILSTRICT?                                 # Apply
    | ID                                                             # Var
    | IF expr THEN expr ( ELSE expr )?                               # IfThenElse
    | op=(PLUS | MINUS | NOT | BITNOT) expr                          # UnaryExpr
    | expr op=(MULTIPLY | DIVIDE | MODULUS) expr                     # BinaryExpr
    | expr op=(PLUS | MINUS) expr                                    # BinaryExpr
    | expr op=(SHIFTLEFT | SHIFTRIGHT) expr                          # BinaryExpr
    | expr op=(GT | GE | LT | LE | IN) expr                          # BinaryExpr
    | expr op=(EQUALS | NOTEQUALS) expr                              # BinaryExpr
    | expr op=BITAND expr                                            # BinaryExpr
    | expr op=BITXOR expr                                            # BinaryExpr
    | expr op=BITOR expr                                             # BinaryExpr
    | expr op=AND expr                                               # BinaryExpr
    | expr op=OR expr                                                # BinaryExpr
    | expr '{' objinside? '}'                                        # ApplyBrace
    | FUNCTION '(' params? ')' expr                                  # Function
    | assertion ';' expr                                             # Assert
    | IMPORT STRING                                                  # Import
    | IMPORTSTR STRING                                               # Import
    | ERROR expr                                                     # ErrorExpr
    | expr IN SUPER                                                  # InSuper
    | LOCAL binds+=bind (',' binds+=bind)* ';' expr                  # LocalBind
    ;


objinside
    : members+=member (',' members+=member)* ','?                    # Members
    | ( objlocal ',' )*  '[' key=expr ']' ':' value=expr ( ',' objlocal )* ','? forspec+
    																 # ObjectComp
    ;

member
    : objlocal | assertion | field
    ;

field
    : fieldname PLUS? visibility expr                                # ValueField
    | fieldname '(' params? ')' visibility expr                      # FunctionField
    ;

visibility
    : ':'
    | ':' ':'
    | ':' ':' ':'
    ;

objlocal
    : LOCAL bind
    ;

forspec
	: FOR ID IN expr ifspec*
	;

ifspec
	: IF expr
	;

fieldname
    : ID
    | STRING
    | '[' expr ']'
    ;

assertion
    : ASSERT condition=expr (':' message=expr)?
    ;


bind
    : ID '=' expr                                                    # ValueBind
    | ID '(' params? ')' '=' expr                                    # FunctionBind
    ;


args
    : pos+=expr ( ',' pos+=expr )* ( ',' names+=ID '=' named+=expr )* ','?
    | names+=ID '=' named+=expr ( ',' names+=ID '=' named+=expr )* ','?
    ;

params
    : pos+=ID ( ',' pos+=ID )* ( ',' names+=ID '=' defaults+=expr )* ','?
    | names+=ID '=' defaults+=expr ( ',' names+=ID '=' defaults+=expr )* ','?
    ;


DOLLAR    : '$';

ASSERT    : 'assert';
ELSE      : 'else';
ERROR     : 'error';
FALSE     : 'false';
FOR       : 'for';
FUNCTION  : 'function';
IF        : 'if';
IMPORT    : 'import';
IMPORTSTR : 'importstr';
LOCAL     : 'local';
NULL      : 'null';
SELF      : 'self';
SUPER     : 'super';
TAILSTRICT: 'tailstrict';
THEN      : 'then';
TRUE      : 'true';

EQUALS    : '==' ;
NOTEQUALS : '!=';
PLUS      : '+';
MINUS     : '-';
MULTIPLY  : '*';
DIVIDE    : '/';
MODULUS   : '%';
AND       : '&&';
OR        : '||';
NOT       : '!';
GT        : '>';
GE        : '>=';
LT        : '<';
LE        : '<=';
IN        : 'in';
SHIFTLEFT : '<<';
SHIFTRIGHT: '>>';
BITNOT    : '~';
BITAND    : '&';
BITXOR    : '^';
BITOR     : '|';

STRING
    : '"' (ESCAPES | UNICODE | ~["\\\u0000-\u001F])* '"'
    | '\'' (ESCAPES | UNICODE | ~['\\\u0000-\u001F])* '\''
    | '@' '"' ('""' | ~["])* '"'
    | '@' '\'' ('\'\'' | ~['])* '\''
    | '@' '\'' ('\'\'' | ~['])* '\''
    | '|||' ( ~'|' | '|' ~'|' |  '||' ~'|' )* '|||'
    ;

NUMBER: INT ( '.' DIGIT+ )? EXP?;

ID: ALPHA (ALPHA | DIGIT)*;

fragment ESCAPES: '\\' ["'\\/bfnrt];
fragment DIGIT: [0-9];
fragment ALPHA: [_a-zA-Z];
fragment UNICODE: 'u' HEX HEX HEX HEX;
fragment HEX: [0-9a-fA-F];
fragment INT: '0' | [1-9] DIGIT*;
fragment EXP: [Ee] [+\-]? DIGIT+;

Whitespace
    : [ \t]+ -> skip
    ;

Newline
    : ( '\r' '\n'? | '\n' ) -> skip
    ;

BlockComment
    : '/*' .*? '*/' -> skip
    ;

LineComment
    : ('//'|'#') ~[\r\n]* -> skip
    ;
