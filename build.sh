!#build/bash
set -e 

CURRENT_BRANCH=$(git branch --show-current)
echo "Current Branch = $CURRENT_BRANCH"

echo "Test build with Go.."
go test./...

echo "Build Docker Image.."
docker build -t devops-lab:v1 .

echo "Build Finished."
