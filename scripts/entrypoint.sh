#!/bin/bash
set -e

# If NGROK_ENABLED is set to true, attempt to fetch the public URL
if [ "$NGROK_ENABLED" = "true" ]; then
    echo "Waiting for ngrok to become available..."
    
    # Wait for ngrok API to be reachable and tunnel to be ready
    echo "Waiting for ngrok tunnel..."
    MAX_RETRIES=30
    COUNTER=0
    while [ $COUNTER -lt $MAX_RETRIES ]; do
        TUNNEL_DATA=$(curl -s http://ngrok:4040/api/tunnels)
        PUBLIC_URL=$(echo "$TUNNEL_DATA" | jq -r '.tunnels[] | select(.proto=="https") | .public_url' | head -n 1)
        
        if [ ! -z "$PUBLIC_URL" ] && [ "$PUBLIC_URL" != "null" ]; then
            echo "Ngrok tunnel found: $PUBLIC_URL"
            export PUBLIC_URL="$PUBLIC_URL"
            break
        fi
        
        echo "Tunnel not ready yet, retrying... ($COUNTER/$MAX_RETRIES)"
        sleep 2
        COUNTER=$((COUNTER+1))
    done

    if [ -z "$PUBLIC_URL" ]; then
        echo "Error: Failed to obtain ngrok URL after timeout."
    fi
fi

# Run the command passed as arguments (or default to the binary)
exec "$@"
