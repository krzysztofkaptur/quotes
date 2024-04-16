!/bin/bash
set -e
 
sed -i "s/uniquevalue/${NEXT_PUBLIC_BASE_API}/g" .env.local
 
exec "$@"