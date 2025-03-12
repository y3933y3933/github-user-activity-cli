# Github User Activity CLI

## 簡介
此專案源自於 roadmap [GitHub User Activity](https://roadmap.sh/projects/github-user-activity) project

## 安裝

1. 安裝 [Go](https://go.dev/)
2. clone 專案
```bash
git clone https://github.com/y3933y3933/github-user-activity-cli
cd github-user-activity-cli
```
3. 建立 .env 檔案到專案根路徑
```bash
API_TOKEN=<Your Github Token>
```

Github Token 產生做法可參考[官網](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-fine-grained-personal-access-token)



## 啟動

```bash
go run . github-activity <username>
```

note:
目前只處理了
- `PushEvent`
- `WatchEvent`

EventType 參考[Github Event Type](https://docs.github.com/en/rest/using-the-rest-api/github-event-types?apiVersion=2022-11-28#issuesevent)
