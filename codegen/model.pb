�
ChatService�

ChatService*�
GET /w
GET /"
patterns
:
"rest:
serve ChatClientB/�8
*dev/socket-sysl/example/specs/example.sysl"%
*|
Connectq
Connect0:)'

ChatServiceChatService -> Connect�8
*dev/socket-sysl/example/specs/example.sysl*r
MOTDj
MOTD0:%#


ChatClientChatService -> MOTD�8
*dev/socket-sysl/example/specs/example.sysl*�
ChatService -> Connect{
ChatService -> Connect*
ChatService:

ChatServiceMOTD�8
*dev/socket-sysl/example/specs/example.sysl*�
ChatClient -> SendMessage�
ChatClient -> SendMessage"
patterns:	
"relay*

ChatClient:


ChatClientSendMessage�8
*dev/socket-sysl/example/specs/example.sysl!�6
*dev/socket-sysl/example/specs/example.sysl�

ChatClient�


ChatClient"
patterns:
"web
"client*y
SendMessagej
SendMessage0:,*

ChatServiceChatClient -> SendMessage:+)


ChatClientChatClient -> SendMessage*M
StartD
Start0�8
*dev/socket-sysl/example/specs/example.sysl((*�
ChatService -> MOTDk
ChatService -> MOTD*
ChatService:

render�8
*dev/socket-sysl/example/specs/example.sysl.2*h
Send`
Send:


ChatClientSendMessage�8
*dev/socket-sysl/example/specs/example.sysl26*�
ChatClient -> SendMessagen
ChatClient -> SendMessage*

ChatClient:

render�6
*dev/socket-sysl/example/specs/example.sysl68�8
*dev/socket-sysl/example/specs/example.sysl%%