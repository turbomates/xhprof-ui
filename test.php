<?php

$serverIp = '127.0.0.1';
$serverPort = 5000;
$message = [
    'data' => '',
    'ID' => '12',
    'project' => 'wst',
    'action' => 'sportbook-worker',
];
$socket = socket_create(AF_INET, SOCK_DGRAM, SOL_UDP);
socket_sendto($socket, json_encode($message), strlen($message), 0, $serverIp, $serverPort);
socket_close($socket);