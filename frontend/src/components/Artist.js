import React, { useState } from 'react';

export default function Artist(
    {
        Name = "Placeholder name",
        URI = "Placeholder uri"
    }
) {
    return (
        <div>
            <p>Name: {Name}</p>
            <p>URI: {URI}</p>
        </div>
    )
}