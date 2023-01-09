#!/bin/bash

fyne bundle icons8-check-mark-button-48.png > bundled_icons.go

fyne bundle -append icons8-check-mark-button-48.png >> bundled_icons.go