<?php

function initCurl()
{
    $ch = curl_init();

    curl_setopt($ch, CURLOPT_URL, "localhost");
    curl_setopt($ch, CURLOPT_PORT, 9200);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

    return $ch;
}

function request($ch) {
    $output = curl_exec($ch);
    curl_close($ch);

    curl_setopt($ch, CURLOPT_HTTPHEADER, [
            'Content-Type: application/json'
        ]
    );

    return $output;
}

function addProduct($ch) {
    $data = json_encode([
        'title' => 'Test',
        'description' => 'Super test description',
        'category' => 1,
    ]);

    curl_setopt($ch, CURLOPT_HTTPHEADER, [
            'Content-Length: ' . strlen($data)]
    );

    curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "PUT");
    curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));

    return request($ch);
}


$ch = initCurl();
echo addProduct($ch);
