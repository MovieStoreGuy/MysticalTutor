# Mystical Tutor - An App to optimise MTG decks
[![GoDoc](https://godoc.org/github.com/RenegadeTech/MysticalTutor?status.svg)](https://godoc.org/github.com/RenegadeTech/MysticalTutor)
[![Maintainability](https://api.codeclimate.com/v1/badges/4581c0b4b4a1ea04a33d/maintainability)](https://codeclimate.com/github/RenegadeTech/MysticalTutor/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/RenegadeTech/MysticalTutor)](https://goreportcard.com/report/github.com/RenegadeTech/MysticalTutor)
[![Build Status](https://travis-ci.org/RenegadeTech/MysticalTutor.svg?branch=master)](https://travis-ci.org/RenegadeTech/MysticalTutor)  

Mystical Tutor allows for any potential brewer to gauge how well there deck will
do and look at making calculated includes for the deck they want to build.
Using this application will aid the brewer and make decisions easier.

## How it works?
This application works on been given a collection of cards to use (cards the brew has available or all cards ever printed) and applying themes and filtering based on the format.

The order of the formats supported are:
- [ ] EDH (Commander)
- [ ] Draft
- [ ] Standard
- [ ] Modern
- [ ] Legacy

Once we have processed your deck list, it evaluates the cards inside your deck and looks to see if there is potential better cards to replace it with.
