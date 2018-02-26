om_api_call() {
  method=$1
  path=$2
  data=$3
  curl -sS -H "Content-Type: application/json" -X $method --fail --digest  -u "$MONGO_OM_USER:$MONGO_OM_API_KEY" ${MONGO_OM_URL}/api/public/v1.0/groups/$MONGO_OM_GROUP_ID/$path -d "$data"
}
