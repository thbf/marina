Marina
======
Crawl specific sites for job postings of data scientists and store and analyze them

Prerequisites
-------------
golang v1.5

Requirements
------------
- Visit google, facebook, linkedin career pages
- Search for term: "data science"
- Follow links and create database of unigrams, bigrams and trigrams
- Create a database of languages and qualifications that companies care about
- Track how many postings per company
- What is the % of data scientist as a % of overall employees
- What are the most desired skills in a data scientist

Installing
----------

    export GOPATH=$PWD/go
    go get golang.org/x/net/html
    export PATH=$PATH:$PWD/go/bin
    go install findlinks

Running
-------

    findlinks https://golang.org
