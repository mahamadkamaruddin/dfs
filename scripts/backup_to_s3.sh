#!/bin/bash
BUCKET_NAME="dfs-storage"
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_FILE="dfs_backup_$TIMESTAMP.tar.gz"

# Create a tarball of the data directory
tar -czvf $BACKUP_FILE data/

# Upload to S3
aws s3 cp $BACKUP_FILE s3://$BUCKET_NAME/backups/

# Delete older backups (keep only the 5 most recent ones)
aws s3 ls s3://$BUCKET_NAME/backups/ | sort | head -n -5 | awk '{print $4}' | while read -r file; do
    aws s3 rm s3://$BUCKET_NAME/backups/"$file"
done

echo "Backup uploaded to S3 bucket: $BUCKET_NAME"
