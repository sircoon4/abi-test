package vm

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sircoon4/bencodex-go"
	"github.com/sircoon4/bencodex-go/util"
)

func TestRun(t *testing.T) {
	serializedPayloadList := [][]byte{}
	serializedPayloadHackAndSlash := []byte("ZDE6UzcwOjBEAiBREB4k6V+8VAXxIe9I1s565xeRyEhTcvft0QtkQavdRAIgF/7gAwEK1hCXBI3vStWHsInR56Pjbpgr77ZGNfYJWJ8xOmFsZHU3OnR5cGVfaWR1MTY6aGFja19hbmRfc2xhc2gyMnU2OnZhbHVlc2R1MTI6YXBTdG9uZUNvdW50dTE6MHUxMzphdmF0YXJBZGRyZXNzMjA63+Olmd3JqvwM5CMsow1BKkN5eE91ODpjb3N0dW1lc2xldTEwOmVxdWlwbWVudHNsMTY6k0hvMSQP3ESqxHe5FvP/ZDE2OnESdTcOEDhHjXX+a8k/kgMxNjputrpDman3Q5V424Y2EJutMTY6vGXnkjiTjk6uEDh/1UnKlTE2OoCmFqSy2FdGlMeyRVDTd0sxNjqAhR6pZsrMRoS+9euOX7TnMTY6yetsy/AeaEm0XJX2GX1KBWV1NTpmb29kc2xldTI6aWQxNjpVMrOqojpcTo+szFsaRu5tdTE6cmxsdTE6MHU1OjMwMDAxZWV1NzpzdGFnZUlkdTM6MjEwdTE0OnRvdGFsUGxheUNvdW50dTE6MXU3OndvcmxkSWR1MTo1ZWVlMTpnMzI6RYIlDQ2jOwZ3moR10oPV3SEMaDubmZ100D+sT1j6a84xOmxpNGUxOm1sZHUxMzpkZWNpbWFsUGxhY2VzMToSdTc6bWludGVyc251Njp0aWNrZXJ1NDpNZWFkZWkxMDAwMDAwMDAwMDAwMDAwMDAwZWUxOm5pMzExNmUxOnA2NToE3wbqFQmA1Y+s/9KKFfFC0scchb6fM0FbN2zIUcbvextqIZ+RGm9zXti+cBRBMhv+/ghXR6k08lAxOUH2iJvMYzE6czIwOmtPF4trn3OiAUK7/onUOjgBB4aDMTp0dTI3OjIwMjQtMDctMDVUMDk6MjM6MjAuNTE3Njc2WjE6dWxlZQ==")
	serializedPayloadGrinding := []byte("ZDE6UzcwOjBEAiBNIJpUkfHyhnA3wCHRFt9iuyHqPbKPbEys9MF/RsG1OQIgTB1dU35a5Cr2odqIkNUE6hEZ/PIdrLt5fiidK3YrQsoxOmFsZHU3OnR5cGVfaWR1OTpncmluZGluZzJ1Njp2YWx1ZXNkdTE6YTIwOlE4aopjoX/XOXBM3WCk5oN/7lWIdTE6Y3R1MTplbDE2OutbnGA0koFAu+zs6J5VuAVldTI6aWQxNjpJtSFKgJLXQpKj9T0cXokPZWVlMTpnMzI6RYIlDQ2jOwZ3moR10oPV3SEMaDubmZ100D+sT1j6a84xOmxpMWUxOm1sZHUxMzpkZWNpbWFsUGxhY2VzMToSdTc6bWludGVyc251Njp0aWNrZXJ1NDpNZWFkZWkxMDAwMDAwMDAwMDAwMDAwMDAwZWUxOm5pMjc1ZTE6cDY1OgTcDJ8HG/Fe9jSl94B4LbBfLyUe2gJZmIVaLbjVMXNKh14btogWXuVmcKGhwxuQeDQHjkVD1NXeLyA+X4Jj5Y0vMTpzMjA6ga9iZHvpgtBiwvnBCbOlOdjJy8ExOnR1Mjc6MjAyNC0wNy0wNVQwOToyMzowNi4xNzYyODhaMTp1bGVl")
	serializedPayloadCombinationEquipment := []byte("ZDE6UzcwOjBEAiAAqENcATSUfl/izYJWShs0nTLf5XDqg+n11hTIlSZ7kAIgWzmye7xj2lTPL1bdO+FztKTgSPHYRXBHTFni5OGjDSMxOmFsZHU3OnR5cGVfaWR1MjM6Y29tYmluYXRpb25fZXF1aXBtZW50MTd1Njp2YWx1ZXNkdTE6YTIwOvsmvr466BLcc1rbsNdF7jN1jxN2dTE6aGZ1MTppdTM6MzczdTI6aWQxNjqwfyDSYCl6QqXK30uDXCoBdTE6cHR1MzpwaWRudTE6cnUxOjF1MTpzdTE6MWVlZTE6ZzMyOkWCJQ0NozsGd5qEddKD1d0hDGg7m5mddNA/rE9Y+mvOMTpsaTFlMTptbGR1MTM6ZGVjaW1hbFBsYWNlczE6EnU3Om1pbnRlcnNudTY6dGlja2VydTQ6TWVhZGVpMTAwMDAwMDAwMDAwMDAwMDAwMGVlMTpuaTQ0ZTE6cDY1OgTtl6zT5sIF42uyL9icLWTkFBfLykCjUDRNSPgVzoJ3P0UPXpBZci4nWTcs/pSwsxcfhgUULLgxb5jhWzvRNMUGMTpzMjA695+PB+EmhdKUge+5CMjRa/kIWmIxOnR1Mjc6MjAyNC0wNy0wNVQwOToyMzoxMy41ODcyMTVaMTp1bGVl")
	serializedPayloadRapidCombination := []byte("ZDE6UzcwOjBEAiBZR4AWgBA3nzr85z78em/zxXYVNUf1FequYqCpyA2JEgIgbsaxuK0+ltCOiXT42GnRxZZYViOrJxM6y5WsNVMh1I8xOmFsZHU3OnR5cGVfaWR1MTk6cmFwaWRfY29tYmluYXRpb24xMHU2OnZhbHVlc2R1MTM6YXZhdGFyQWRkcmVzczIwOvH7GPXpy7CHQO4PFtgXbP3CgQNxdTI6aWQxNjoq3OyMSPUsS5aa18gH++MtdTk6c2xvdEluZGV4dTE6MGVlZTE6ZzMyOkWCJQ0NozsGd5qEddKD1d0hDGg7m5mddNA/rE9Y+mvOMTpsaTFlMTptbGR1MTM6ZGVjaW1hbFBsYWNlczE6EnU3Om1pbnRlcnNudTY6dGlja2VydTQ6TWVhZGVpMTAwMDAwMDAwMDAwMDAwMDAwMGVlMTpuaTE2M2UxOnA2NToEU/MbuCWFARsi1/TWHi1mmXKhjO7X1C0T7Iy81RiItB3qBfawFHS3Tn0asS1hW8XHstN1mLKixzKRIutHO9vbCDE6czIwOttjcXDqIhLYPQyav09/1q7lFQgAMTp0dTI3OjIwMjQtMDctMDVUMDk6MjM6MDcuNzkzODE2WjE6dWxlZQ==")
	serializedPayloadHackAndSlashSweep := []byte("ZDE6UzcxOjBFAiEA3pSeUIw3Q7ZzkkeB4Iqa1d/W8OVdS55nCqJksCw40DYCIDCByMHbJqzes+gCsQiH0JzZ5zl3FN1bjOnwwYoT4G39MTphbGR1Nzp0eXBlX2lkdTIyOmhhY2tfYW5kX3NsYXNoX3N3ZWVwMTB1Njp2YWx1ZXNkdTExOmFjdGlvblBvaW50dTM6MTE1dTEyOmFwU3RvbmVDb3VudHUxOjB1MTM6YXZhdGFyQWRkcmVzczIwOgpwc9j+o4OCu1j37+QK3DU8CMMsdTg6Y29zdHVtZXNsZXUxMDplcXVpcG1lbnRzbDE2OvFS0g6N6xZNq1uSRRFzhfoxNjrPDwwqGNsyTqDLCCpn6R0xMTY6d3MmZy76p065VNQ9D8I2yTE2OvPoHX2fFWRIuV85tK9wzPsxNjpxKy6HMuBATo/kcggaW2NbMTY6UGCdqF7Fe0OlZEm2wm7rUzE2OuugcKljm7ZJrtOd8J5oqzVldTI6aWQxNjqS2IHQ4BQaTKKONtIVVwQgdTk6cnVuZUluZm9zbGx1MTowdTU6MzAwMDFlZXU3OnN0YWdlSWR1MzoxMzl1Nzp3b3JsZElkdTE6M2VlZTE6ZzMyOkWCJQ0NozsGd5qEddKD1d0hDGg7m5mddNA/rE9Y+mvOMTpsaTFlMTptbGR1MTM6ZGVjaW1hbFBsYWNlczE6EnU3Om1pbnRlcnNudTY6dGlja2VydTQ6TWVhZGVpMTAwMDAwMDAwMDAwMDAwMDAwMGVlMTpuaTI1N2UxOnA2NToEFNrx5bPvcPh05jtabwKOrsbvGhukSlWNA2ydhLJ6/Kvb07VgrHTQJ3CuX9KjKBkGrW/E3EMnpe3QkSVO8RSNxjE6czIwOveKSbie1JodVw+wZeAskPe1SH5bMTp0dTI3OjIwMjQtMDctMDVUMDk6MjM6MTIuNTg3MDAyWjE6dWxlZQ==")
	serializedPayloadTransferAsset := []byte("ZDE6UzcwOjBEAiBSOidioUVetuLwDfe/GPEoIK57gJ9eCjClsrzgwJOvmgIgHvRYjATEzc3qL9f/a4bo1i/2iHwcyAiB3w9EnbkKvvQxOmFsZHU3OnR5cGVfaWR1MTU6dHJhbnNmZXJfYXNzZXQ1dTY6dmFsdWVzZHU2OmFtb3VudGxkdTEzOmRlY2ltYWxQbGFjZXMxOgJ1NzptaW50ZXJzbDIwOkfQgqEVxj57WLFTLSDmMVOOr63eZXU2OnRpY2tlcnUzOk5DR2VpMTEwNGVldTk6cmVjaXBpZW50MjA6gjF6uyNzFLab6VM+b/TVpIsdG5J1NjpzZW5kZXIyMDo8gmHRKY1szqMo5jG2hFJCtgBxb2VlZTE6ZzMyOkWCJQ0NozsGd5qEddKD1d0hDGg7m5mddNA/rE9Y+mvOMTpsaTRlMTptbGR1MTM6ZGVjaW1hbFBsYWNlczE6EnU3Om1pbnRlcnNudTY6dGlja2VydTQ6TWVhZGVpMTAwMDAwMDAwMDAwMDAwMDAwMGVlMTpuaTQ4MWUxOnA2NToEJz5TlBFQ9EUwolEUJxDrQg9N55WqoP/N3yC85UouHhFChRgZ0v60DJxw84lyvsXntIcSLQq6iNzwWcO5bHDLczE6czIwOjyCYdEpjWzOoyjmMbaEUkK2AHFvMTp0dTI3OjIwMjQtMDctMDVUMDk6MjM6MjEuMTUwMzY1WjE6dWxlZQ==")
	serializedPayloadClaimItems := []byte("ZDE6UzcxOjBFAiEAs6Dh4AKfGoj/JXt3105Q1AOLaE7BLTb5jAWhYc8DEyQCIAFualk5WtoTWT2mweMLgZRDbvkAG6KWdb3kt+UZRnunMTphbGR1Nzp0eXBlX2lkdTExOmNsYWltX2l0ZW1zdTY6dmFsdWVzZHUyOmNkbGwyMDprqXvayCAEA/BBF/4us4k6Z2/RomxsZHUxMzpkZWNpbWFsUGxhY2VzMToSdTc6bWludGVyc251Njp0aWNrZXJ1MTI6RkFWX19DUllTVEFMZWkyNDAwMDAwMDAwMDAwMDAwMDAwMDAwZWVsZHUxMzpkZWNpbWFsUGxhY2VzMToAdTc6bWludGVyc251Njp0aWNrZXJ1MTQ6SXRlbV9OVF84MDAyMDFlaTEyZWVlZWV1MjppZDE2Oh380p2eg9BFhvxr+YVClCl1MTptdTYwOnBhdHJvbCByZXdhcmQgNmJhOTdCREFjODIwMDQwM2YwNDExN0ZFMmVCMzg5M2E2NzZmRDFBMiAvIDEwN2VlZTE6ZzMyOkWCJQ0NozsGd5qEddKD1d0hDGg7m5mddNA/rE9Y+mvOMTpsaTRlMTptbGR1MTM6ZGVjaW1hbFBsYWNlczE6EnU3Om1pbnRlcnNudTY6dGlja2VydTQ6TWVhZGVpMTAwMDAwMDAwMDAwMDAwMDAwMGVlMTpuaTIzODYwNDhlMTpwNjU6BI+GguV92N77nV3VKLxTstLYZ6TAt7bz+YtvRKFnhSUmzMZqbTUSvEKcIOICgBUry/7opQPAKMVtXzSLYkidLnwxOnMyMDrK1g8YtLoYn38cFOImfZsg9bFv9TE6dHUyNzoyMDI0LTA3LTA2VDA5OjIzOjI0LjAwMjAxNloxOnVsZWU=")
	serializedPayloadDailyReward := []byte("ZDE6UzcxOjBFAiEAn3VGdyP4bKNy33mHx3uxqzA0734n02xYCWk9gK4H9LMCID07fnplsCNx3bTLjWPhkAGtB/IdjYA8dvYEZdkWyD9pMTphbGR1Nzp0eXBlX2lkdTEzOmRhaWx5X3Jld2FyZDd1Njp2YWx1ZXNkdTE6YTIwOlV3ZRo0zm7Z7BcWrho/wGGXFPaGdTI6aWQxNjpcuHt1loAnS7tZ9UzuDaMzZWVlMTpnMzI6RYIlDQ2jOwZ3moR10oPV3SEMaDubmZ100D+sT1j6a84xOmxpMWUxOm1sZHUxMzpkZWNpbWFsUGxhY2VzMToSdTc6bWludGVyc251Njp0aWNrZXJ1NDpNZWFkZWkxMDAwMDAwMDAwMDAwMDAwMDAwZWUxOm5pNDc2ZTE6cDY1OgTaEBUYMLJrA9mBfKWJFo7NEG+q6M5mROwfI/slkXTBoF2i7ceV8zibURlLo6XLfQ7Ywuxw9Y3f+6Y8dyIll0zNMTpzMjA61bNT7+Q6Rn9QXME3gLfr2axHB2wxOnR1Mjc6MjAyNC0wNy0wOVQwMToyMzo1Ni45ODQ5NzNaMTp1bGVl")
	serializedPayloadAuraSummon := []byte("ZDE6UzcwOjBEAiBJgn0kAWoP9DGI5cboIE6Qjvcdn2Em5WtOFu26jGthBgIgAs16li3KifglJqFgLO413y/p6J6EqBkbWERiJbCOsZIxOmFsZHU3OnR5cGVfaWR1MTE6YXVyYV9zdW1tb251Njp2YWx1ZXNkdTI6YWEyMDphJnpp1xh/lUDLGy3ZGIF91nq4jnUzOmdpZHU1OjEwMDAxdTI6aWQxNjrLNz3HOubmSK/G+uSeO2mZdTI6c2N1MToxZWVlMTpnMzI6RYIlDQ2jOwZ3moR10oPV3SEMaDubmZ100D+sT1j6a84xOmxpMWUxOm1sZHUxMzpkZWNpbWFsUGxhY2VzMToSdTc6bWludGVyc251Njp0aWNrZXJ1NDpNZWFkZWkxMDAwMDAwMDAwMDAwMDAwMDAwZWUxOm5pMTdlMTpwNjU6BMTgoG/+EmpSa21Wnpd0u8Zy19MiJV0BVWH0muMOjfTDBOlpROqNwJ+6CCoONS6gJp3cVjdmKet8ViIB1INrsisxOnMyMDq6ALc2ZrfVKaXqGwY7uG0C8kN3szE6dHUyNzoyMDI0LTA3LTA5VDAyOjE3OjU0Ljg0NTg0NFoxOnVsZWU=")
	serializedPayloadExploreAdventureBoss := []byte("ZDE6UzcxOjBFAiEAyvaMltM6Yd7M2tXoKi2svYjIyIIrHb3ZOuJp2C54Ro8CIGvLL6DgODLYS9kKB5AAB3rhr29WeB2gigqdIJQEPbhHMTphbGR1Nzp0eXBlX2lkdTIyOmV4cGxvcmVfYWR2ZW50dXJlX2Jvc3N1Njp2YWx1ZXNkdTEzOmF2YXRhckFkZHJlc3MyMDr1lcjRhouvmYfv9u83t3+dXnx76HU4OmNvc3R1bWVzbDE2OmzzZyAOikZGia7o7G9L5/8xNjrXc8pzP5riSZWC3YGIHHVEMTY6SxKCmAO+jUuqAwRCtvFQimV1MTA6ZXF1aXBtZW50c2wxNjptYhIDCAYYQ5UEmJvcBnbvMTY6nMZjNN/GQkWZ1s0/Fd210DE2OpblZ14n4ftJtFVmVN5dy50xNjrRsFBhn/ZSSakCf20/xgZsMTY6i2TJcUjDgk+xYyeTPVQ6YjE2Ory0ZrIkXx1Ao7abFqLckqExNjrWBLPU40NsQLGt4kdqSpqwZXU1OmZvb2RzbGV1MjppZDE2Oi5XDqIwys9Ol9QOFhbsPiR1MTpybGx1MTowdTU6MTAwMTJlbHUxOjF1NToxMDAxMWVsdTE6M3U1OjEwMDI5ZWx1MTo0dTU6MTAwMDNlbHUxOjZ1NToxMDAwMmVsdTE6N3U1OjIwMDAxZWV1NjpzZWFzb25pMWVlZWUxOmczMjpFgiUNDaM7BneahHXSg9XdIQxoO5uZnXTQP6xPWPprzjE6bGkxZTE6bWxkdTEzOmRlY2ltYWxQbGFjZXMxOhJ1NzptaW50ZXJzbnU2OnRpY2tlcnU0Ok1lYWRlaTEwMDAwMDAwMDAwMDAwMDAwMDBlZTE6bmkyNDg0ZTE6cDY1OgRGb/NIPk7XOh3sKyt7ybABoz7vlYE6pagZXcBesbBcLU3pn2noeE+gqyyGUVaGa37OQ6QskzRnhYnb7ETmcZUGMTpzMjA6B4QmPkZhC1rqpBDthNmX1Q47iDMxOnR1Mjc6MjAyNC0wNy0wOVQwMjoxODoxMC4zMDk3OTVaMTp1bGVl")

	serializedPayloadList = append(serializedPayloadList, serializedPayloadHackAndSlash)         //0
	serializedPayloadList = append(serializedPayloadList, serializedPayloadGrinding)             //1
	serializedPayloadList = append(serializedPayloadList, serializedPayloadCombinationEquipment) //2
	serializedPayloadList = append(serializedPayloadList, serializedPayloadRapidCombination)     //3
	serializedPayloadList = append(serializedPayloadList, serializedPayloadHackAndSlashSweep)    //4
	serializedPayloadList = append(serializedPayloadList, serializedPayloadTransferAsset)        //5
	serializedPayloadList = append(serializedPayloadList, serializedPayloadClaimItems)           //6
	serializedPayloadList = append(serializedPayloadList, serializedPayloadDailyReward)          //7
	serializedPayloadList = append(serializedPayloadList, serializedPayloadAuraSummon)           //8
	serializedPayloadList = append(serializedPayloadList, serializedPayloadExploreAdventureBoss) //9

	for i, serializedPayload := range serializedPayloadList {
		fmt.Println("TestRun() i:", i)
		_ = i
		abi, err := Run(serializedPayload)
		if err != nil {
			t.Errorf("%dth Run() failed: %v", i, err)
		}
		_ = abi
		fmt.Println("abi:", abi)
		fmt.Println()
	}
}

func TestRunFromRealDatas(t *testing.T) {
	const dirPath = "./error_datas"
	const spFilePath = "./error_datas/serialized_payloads_%s_%d.dat"
	const jsonFilePath = "./error_datas/json_%s_%d.json"

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	path9c := os.Getenv("9C_GRAPHQL_EXPLORER_API_URL")

	// Make GraphQL query request
	query := `{
		blockQuery{
			blocks(desc: true, limit: 1) {
				transactions {
					serializedPayload
				}
			}
		}
	}`

	// Create the request body
	requestBody, err := json.Marshal(map[string]string{
		"query": query,
	})
	if err != nil {
		fmt.Println("Error creating request body:", err)
		return
	}

	// Send the request
	resp, err := http.Post(path9c, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Unmarshal the response body
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response body:", err)
		return
	}

	var serializedPayloadList [][]byte
	for _, transaction := range response.Data.BlockQuery.Blocks[0].Transactions {
		serializedPayloadList = append(serializedPayloadList, []byte(transaction.SerializedPayload))
	}

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	for i, serializedPayload := range serializedPayloadList {
		_, err := Run(serializedPayload)
		if err != nil {
			t.Errorf("%dth Run() failed: %v", i, err)
			errStr := string(err.Error())

			err = os.WriteFile(fmt.Sprintf(spFilePath, errStr, i), []byte(serializedPayload), 0644)
			if err != nil {
				fmt.Println("Error writing serializedPayload data:", err)
				return
			}

			serializedPayloadEncoded, err := base64.StdEncoding.DecodeString(string(serializedPayload))
			if err != nil {
				fmt.Println("Error decoding serialized payload:", err)
				return
			}
			bencodexValue, err := bencodex.Decode(serializedPayloadEncoded)
			if err != nil {
				fmt.Println("Error decoding bencodex:", err)
				return
			}
			bencodexJson, err := util.MarshalJsonRepr(bencodexValue)
			if err != nil {
				fmt.Println("Error marshalling json:", err)
				return
			}
			err = os.WriteFile(fmt.Sprintf(jsonFilePath, errStr, i), bencodexJson, 0644)
			if err != nil {
				fmt.Println("Error writing Bencodex json data:", err)
				return
			}
		}
	}
}
