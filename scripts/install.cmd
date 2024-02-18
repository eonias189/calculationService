cd orchestrator && go mod download && cd ..
cd backend && go mod download && cd ..
cd agent && go mod download && cd ..
cd frontend && yarn install && yarn build && cd ..