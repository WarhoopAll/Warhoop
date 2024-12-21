#!/bin/sh

COMMIT_HASH=$(git rev-parse HEAD 2>/dev/null || echo "N/A")
BRANCH=${1:-$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "N/A")}
COMMIT_DATE=$(git log -1 --format=%cd --date=format:"%Y-%m-%d %H:%M:%S" 2>/dev/null || echo "N/A")
LAST_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "N/A")

echo "Commit Hash: $COMMIT_HASH"
echo "Branch: $BRANCH"
echo "Commit Date: $COMMIT_DATE"
echo "Last Tag: $LAST_TAG"

COMMITS_RAW=$(git log --all --pretty=format:'{"date": "%ad", "author": "%an", "message": "%s", "commit": "%H"}' --date=format:"%Y-%m-%d" 2>/dev/null || echo "[]")

echo "DEBUG: Raw Commits:"
echo "$COMMITS_RAW"

COMMITS_RAW=$(echo "$COMMITS_RAW" | jq -R -s '
  [splits("\n") | select(length > 0) | fromjson?]')

echo "DEBUG: Processed Raw Commits:"
echo "$COMMITS_RAW" | jq .

TAGS_RAW=$(git tag -l --sort=-version:refname | xargs -I {} git log -1 --format='{"tag": "{}", "date": "%ad"}' --date=format:"%Y-%m-%d" {} 2>/dev/null || echo "[]")

TAGS=$(echo "$TAGS_RAW" | jq -R -s '
  [splits("\n") | select(length > 0) | fromjson?]')

echo "DEBUG: Processed Tags:"
echo "$TAGS" | jq .

COMMITS=$(echo "$COMMITS_RAW" | jq --argjson tags "$TAGS" '
  group_by(.date) |
  map({
    day: .[0].date,
    tag: "none",
    commits: map({
      date: .date,
      author: .author,
      message: .message
    })
  }) |
  reduce .[] as $commit ([];
    . + [$commit | .tag = (
      ($tags | map(select(.date <= $commit.day)) | sort_by(.date) | last // {tag: "none"}).tag
    )]
  ) |
  sort_by(.day) | reverse
')

echo "DEBUG: Processed Commits:"
echo "$COMMITS" | jq .

AUTHORS_RAW=$(git shortlog -sn --all 2>/dev/null || echo "")
AUTHORS=$(echo "$AUTHORS_RAW" | awk '
BEGIN { authors_list = ""; }
{
  commits=$1;
  $1="";
  name=substr($0, index($0,$2));
  authors[name]+=commits;
}
END {
  print "[";
  first=1;
  for (author in authors) {
    if (!first) {
      print ",";
    }
    first=0;
    print "  {\"name\": \"" author "\", \"commits\": " authors[author] "}";
  }
  print "]";
}')

AUTHORS=$(echo "$AUTHORS" | sed '$ s/,$//')

echo "Processed Authors:"
echo "$AUTHORS"

cat <<EOF > ./static/gitinfo.json
{
  "commitHash": "$COMMIT_HASH",
  "branch": "$BRANCH",
  "commitDate": "$COMMIT_DATE",
  "lastTag": "$LAST_TAG",
  "authors": $AUTHORS,
  "commits": $COMMITS
}
EOF

echo "Git info saved to gitinfo.json"