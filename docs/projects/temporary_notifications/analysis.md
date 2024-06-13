# Analysis

## Introduction

What could be useful in all kind of applications are temporary notifications.

Temporary because the notification will automatically disapear after the given timeout.
Notification because the message could either be an error, information, warning or anything else.

This approach avoid to lock the application behind a modal error dialog.

In many cases we would prefer a dynamic approach.

## Concept

I want to design it as a widget that could have actions.

Like this : 

`gNotification.Error("message") // Shows an Error notification in the widget with the given message.`

The configuration will happen at the initialization of the widget. Like this :

`gNotification.SetNotificationTimer(5) // Sets the timer to 5 sec.`

The thing is that I need a queue to stack incomming notifications.

I think, the expected behaviour would be that the notifications are show on top of each others,
and gradually disapears when their time is over.

Maybe notifications should be widgets but the system managing them should be something more 'global'.

I need to confirm, but I think that I could place widget with absolute positions. Tho I don't think
that I can display widgets without having a layout or at least a container.
According to the Fyne's docs I need a container without layout to be able to show items where I want.

So maybe the app should be a stack with the actual content of the app and a container that Komok manages.

For exemple : 

`window.SetContent(container.NewStack(makeGui, komok.content))`

The problem would be to manage possibly multiple window...

And it would be good that those notifications can have some kind of smooth animations...

