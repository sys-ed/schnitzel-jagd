Given that both server and client run on the same machine we
should be able to find some information.

In this case we are using UDP. We can look at

$ cat /proc/net/udp
  sl  local_address rem_address   st tx_queue rx_queue tr tm->when retrnsmt   uid  timeout inode ref pointer drops             
  497: XXXXXXXX:YYYY 00000000:0000 07 00000000:00000000 00:00000000 00000000     0        0 50373 2 0000000000000000 0         
  508: XXXXXXXX:YYYY 00000000:0000 07 00000000:00000000 00:00000000 00000000     0        0 49889 2 0000000000000000 0 


$ ss -lu
State             Recv-Q            Send-Q                       Local Address:Port                         Peer Address:Port            Process            
UNCONN            0                 0                              A.B.C.D:YYYY                             0.0.0.0:*                                  
UNCONN            0                 0                               A.B.C.D:XXXX                              0.0.0.0:*  

Or the more classic netstat utility. Notice how we have a listening datagram
socket for both the server and the client.
