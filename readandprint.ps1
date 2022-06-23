$npipeserver = New-Object System.IO.Pipes.NamedPipeServerStream('pipe1', [System.IO.Pipes.PipeDirection]::InOut)
$npipeserver.WaitForConnection()
"Connection"
$pipeReader = New-Object System.IO.StreamReader($npipeserver)

$line = $pipeReader.ReadLine()
$line
"This line is read: $line"

Start-Sleep -Seconds 3

$npipeserver.Dispose()