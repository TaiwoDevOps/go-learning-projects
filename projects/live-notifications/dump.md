can there be multiple live notifications for different transactions and is this a best practice. 

User sends money to multiple receivers at almost the same time, not needed for the PoC but asking for clarity. 

Can this system work without using APNS/FCM but a websocket ? 

For this PoC which is really just a weekend job, can we strip it down to barest minimum, eg, removing the token management, this application just needs to run a single device that i can record its screen. 

Using a transaction model doesn't necessarily mean the full mechanics, eg, no need for reconciliation handling. If the PoC is what is blowing it up, limit it to just MVP. 
