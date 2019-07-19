 # google cloud sdk configuration

curl https://sdk.cloud.google.com | bash > /dev/null;
source $HOME/google-cloud-sdk/path.bash.inc
gcloud components update kubectl

echo ${GOOGLE_AUTH} > ${HOME}/service-account.json
gcloud auth activate-service-account --key-file ${HOME}/service-account.json
gcloud config set project sublime-delight-246503
gcloud config set compute/zone us-east1-b
gcloud container clusters get-credentials micro-cluster 

kubectl apply -f k8s

