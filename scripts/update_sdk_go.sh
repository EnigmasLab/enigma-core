# On Mac OS: brew install gnu-sed
# On Linux: change command to sed

cp enigma-chain/modules/auction/types/*.go ../sdk-go/chain/auction/types/
cp enigma-chain/modules/exchange/types/*.go ../sdk-go/chain/exchange/types/
cp enigma-chain/modules/ocr/types/*.go ../sdk-go/chain/ocr/types/
cp enigma-chain/modules/peggy/types/*.go ../sdk-go/chain/peggy/types/
cp enigma-chain/modules/wasmx/types/*.go ../sdk-go/chain/wasmx/types/
cp enigma-chain/modules/insurance/types/*.go ../sdk-go/chain/insurance/types/
cp enigma-chain/modules/oracle/types/*.go ../sdk-go/chain/oracle/types/
cp enigma-chain/modules/tokenfactory/types/*.go ../sdk-go/chain/tokenfactory/types/
cp -r proto/ ../sdk-go/proto


cd ../sdk-go/chain/auction/types/
rm -f *test.go
rm -f *gw.go
gsed -i "s/github.com\/EnigmasLab\/enigma-core\/enigma-/github.com\/EnigmasLab\/sdk-go\//g" *.go

cd ../../exchange/types/
rm -f *test.go
rm -f *gw.go
gsed -i "s/github.com\/EnigmasLab\/enigma-core\/enigma-chain\/modules/github.com\/EnigmasLab\/sdk-go\/chain/g" *.go

cd ../../ocr/types/
rm -f *test.go
rm -f *gw.go
gsed -i "s/github.com\/EnigmasLab\/enigma-core\/enigma-chain\/modules/github.com\/EnigmasLab\/sdk-go\/chain/g" *.go

cd ../../peggy/types/
rm -f *test.go
rm -f *gw.go
gsed -i "s/github.com\/EnigmasLab\/enigma-core\/enigma-chain\/modules/github.com\/EnigmasLab\/sdk-go\/chain/g" *.go

cd ../../wasmx/types/
rm -f *test.go
rm -f *gw.go
gsed -i "s/github.com\/EnigmasLab\/enigma-core\/enigma-chain\/modules/github.com\/EnigmasLab\/sdk-go\/chain/g" *.go

cd ../../insurance/types/
rm -f *test.go
rm -f *gw.go
gsed -i "s/github.com\/EnigmasLab\/enigma-core\/enigma-chain\/modules/github.com\/EnigmasLab\/sdk-go\/chain/g" *.go

cd ../../oracle/types/
rm -f *test.go
rm -f *gw.go
gsed -i "s/github.com\/EnigmasLab\/enigma-core\/enigma-chain\/modules/github.com\/EnigmasLab\/sdk-go\/chain/g" *.go

cd ../../tokenfactory/types/
rm -f *test.go
rm -f *gw.go
gsed -i "s/github.com\/EnigmasLab\/enigma-core\/enigma-chain\/modules/github.com\/EnigmasLab\/sdk-go\/chain/g" *.go
