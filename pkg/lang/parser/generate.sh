#!/bin/sh
alias antlr4='java -Xmx500M -cp "./../../../antlr.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -visitor -no-listener -package parser *.g4
