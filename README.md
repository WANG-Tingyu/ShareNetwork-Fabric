  Backend: https://github.com/WANG-Tingyu/ShareNetwork-Fabric.git
  Frontend: https://github.com/WANG-Tingyu/txfabric-app.git
  Demo: https://youtu.be/K3EcpkUQH8I


Network Topology

Three Orgs(Peer Orgs)

    - Each Org have one peer(Each Endorsing Peer)
    - Each Org have separate Certificate Authority
    - Each Peer has Current State database as couch db


One Orderer Org

    - Three Orderers
    - One Certificate Authority



Steps:

1) Clone the repo
2) Run Certificates Authority Services for all Orgs
3) Create Cryptomaterials for all organizations
4) Create Channel Artifacts using Org MSP
5) Create Channel and join peers
6) Deploy Chaincode
   1) Install All dependency
   2) Package Chaincode
   3) Install Chaincode on all Endorsing Peer
   4) Approve Chaincode as per Lifecycle Endorsment Policy
   5) Commit Chaincode Defination
7) Create Connection Profiles
8) Start API Server
9) Register User using API
10) Invoke Chaincode Transaction
11) Query Chaincode Transaction


cd FabricNetwork-2.x/

cd artifacts/channel/create-certificate-with-ca
docker-compose up -d
docker ps
./create-certificate-with-ca.sh


cd ..
./create-artifacts.sh

cd ..
docker-compose up -d
docker ps

cd ..
./createChannel.sh

./deployTransactionCC.sh
postman
UI启动


换fabcar演示

cd artifacts/src/github.com/fabcar/go/
go mod tidy
cd ../../../../..

./depolyChaincode.sh
(Show agreement from those orgs)

cd api-2.0/
npm install

cd config/
./generate-ccp.sh

cd ..
nodemon app.js 

(如果遇到4000端口占用：
sudo lsof -i :4000

sudo kill <PID>
)

成功启用，回到本机，关联端口
ssh -L 4000:127.0.0.1:4000  

打开postman
reg user, get returned token
Paste token to add car

去couchdb查看新添加的，回到本机terminal，关联端口
ssh -L 5984:127.0.0.1:5984 
couchdb UI: http://localhost:5984/_utils/#login
Username: admin
Password: adminpw

Back to postman, add car, keep same car_id, modify others
Go to couchDB, find that record. rev, which represents version change to 2.

Back to postman, find getCarByID. Paste user token. Then call it.

Deploy new Chaincode:
Back to server, open new terminal in FabricNetwork-2.x
./deployDocumentCC.sh 

After deploy, show 3 more containers about documentCC
docker ps

Deploy explorer:
First, copy crypto-config folder to Explorer:
cp -r artifacts/channel/crypto-config Explorer/

/home/tingyu/myFabric/FabricNetwork-2.x/Explorer/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp/keystore
Rename it to 'priv_sk'

cd Explorer/
docker-compose up -d
 
