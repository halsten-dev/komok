# Analysis

## Introduction

The fact is : I want to manage a "Recently opened projects" in DevTracker.

For this I need to manage child menu :

File
    Recent project
        Project#1
        Project#2
        etc...

Actually my menu manager is far too rigid. I need something more flexible.

## The idea

The idea I have is to work only with ID and autoincrement order.

So basically, the manager will hold a list of menu ID and the corresponding data.

A menu can be linked to a menuItem ID.

Then the menu will be built by going through the list and declaring everything as needed.

But the hardest part in this, is that the childmenu needs to be dynamic and updatable.

So the item list of this menu will need constant update. Everytime we open a project in DevTracker
the menu should be updated.