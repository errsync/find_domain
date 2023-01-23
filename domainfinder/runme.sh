#!/bin/bash
echo Buduję program domainfinder...
go build -o domainfinder.exe
echo Buduję program synonyms...
cd ../synonyms
go build -o ../domainfinder/lib/synonyms.exe
echo Buduję program available...
cd ../available
go build -o ../domainfinder/lib/available.exe
echo Buduję program sprinkle...
cd ../sprinkle
go build -o ../domainfinder/lib/sprinkle.exe
echo Buduję program coolify...
cd ../coolify
go build -o ../domainfinder/lib/coolify.exe
echo Buduję program domainify...
cd ../domainify
go build -o ../domainfinder/lib/domainify.exe
echo Budowanie zestawu zakończone.

# Po utworzeniu tego pliku nalezy zmienic dostep (chmod -x runme.sh)

# skrytp buduje programy do folderu ../domainfinfer/lib, nasętpnie należy odpalić program główny domainfinder