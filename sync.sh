echo "Generating hugo"
hugo
echo "Syncing public/ with thj-wiki.web.app..."
aws s3 sync public/ s3://thj-wiki.web.app --acl public-read --region us-west-2 --delete --profile xackery