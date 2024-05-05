# actions_notification

## latest version
anyappinc/actions-discord@v0.0.1

## usage

```yml
name: actions notification test
on: [pull_request, push]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
        with:
          go-version: 1.22

      - name: Run Code
        run: go run main.go

      - name: Hoge
        run: something

      # you need to set environment variables
      - name: Setup environment
        if: Always()
        run: |
          echo "WEBHOOK_URL=${{ secrets.WEBHOOK_URL }}" >> $GITHUB_ENV
          echo "GITHUB_STATUS=${{ job.status }}" >> $GITHUB_ENV
          echo "GITHUB_COMMIT_MESSAGE=${{ github.event.commits[0].message }}" >> $GITHUB_ENV
          echo "GITHUB_PR_TITLE=${{ github.event.pull_request.title }}" >> $GITHUB_ENV
          echo "GITHUB_PR_URL=${{ github.event.pull_request.html_url }}" >> $GITHUB_ENV

      # set notification step
      - name: Notify
        if: Always()
        uses: anyappinc/actions-discord@v0.0.1
```

## Environment variables

you need to set these environment variables in your yml file.

```yml
echo "WEBHOOK_URL=${{ secrets.WEBHOOK_URL }}" >> $GITHUB_ENV
echo "GITHUB_STATUS=${{ job.status }}" >> $GITHUB_ENV
echo "GITHUB_COMMIT_MESSAGE=${{ github.event.commits[0].message }}" >> $GITHUB_ENV
echo "GITHUB_PR_TITLE=${{ github.event.pull_request.title }}" >> $GITHUB_ENV
echo "GITHUB_PR_URL=${{ github.event.pull_request.html_url }}" >> $GITHUB_ENV

```
