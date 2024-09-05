# Email Sequence Management API

This project is an API designed to manage email sequences and their steps, allowing you to create, retrieve, and update sequences with tracking capabilities. The API helps manage the configuration of email sequences, their associated steps, and tracking options for open and click tracking.

## Features

- **Sequences**: Each sequence can contain multiple steps, each with its own subject and content.
- **Tracking Options**: Open and click tracking can be enabled or disabled for each sequence.
- **Sequence Steps**: Steps within a sequence are managed in an ordered way, specifying the content of each email to be sent.
- **API Endpoints**: The API provides several endpoints for managing sequences, retrieving them, and updating tracking options.

## Database Schema

### `sequences`
| Column                  | Type        | Description                                 |
|-------------------------|-------------|---------------------------------------------|
| id                      | SERIAL      | Primary key                                 |
| name                    | VARCHAR(255)| Name of the sequence                        |
| open_tracking_enabled    | BOOLEAN     | If open tracking is enabled for the sequence|
| click_tracking_enabled   | BOOLEAN     | If click tracking is enabled for the sequence|
| wait_hours               | INT         | The waiting time between sending steps      |

### `sequence_steps`
| Column       | Type        | Description                                 |
|--------------|-------------|---------------------------------------------|
| id           | SERIAL      | Primary key                                 |
| sequence_id  | INT         | Foreign key referencing `sequences`         |
| subject      | VARCHAR(255)| Email subject                               |
| content      | TEXT        | Email content                               |
| step_order        | INT         | Step order in the sequence                  |

## API Endpoints

### Create Sequence

Creates a new email sequence with the provided steps and settings.

**POST** `/sequences`

**Request Payload:**
```json
{
  "name": "New Sequence",
  "open_tracking_enabled": true,
  "click_tracking_enabled": false,
  "steps": [
    {
      "subject": "Welcome Email",
      "content": "Welcome to our service!"
    },
    {
      "subject": "Follow-up Email",
      "content": "We hope you're enjoying our service."
    }
  ]
}
```


### Update Sequence Tracking
**PATCH** `/sequences/{id}/tracking`

Updates tracking options for a sequence, such as enabling/disabling open and click tracking.

**Request Payload:**
```json
{
  "open_tracking_enabled": false,
  "click_tracking_enabled": true
}
```

### Get Sequence by ID
**GET** `/sequences/{id}`

Retrieves a specific email sequence along with its steps.

### Get All Sequences
**GET** `/sequences`

Fetches all sequences and their associated steps.

### Update step by sequence and step id
**PUT** `/sequences/{id}/steps/{stepId}`

Updates Subject and/or Content of step.

**Request Payload:**
```json
{
  "subject": "new subject",
  "content": "new content"
}
```

### Delete Step by sequence and step id
**DELETE** `/sequences/{id}/steps/{stepId}`

Deletes a step of a sequence.

## Running the Project

To run this project, ensure you have Docker and Docker Compose installed.

### Steps:

1. Clone the repository:
    ```bash
    git clone <repo-url>
    cd <repo-directory>
    ```

2. Build and run the containers:
    ```bash
    make all
    ```

3. Access the API at `http://localhost:8080`.

## Testing

To run the unit tests:
```bash
make test
```

## Further work

Additionally to this API there's a plan of adding cronjob to manage the distribution of emails in sequences, while respecting the daily email limits of associated mailboxes. The cronjob will handle email delivery to contacts and ensures emails are sent at equal intervals, while considering step wait times and available sending capacities.

Ideally should be hosted in AWS Lambda.

[Here's the flowchart for the cronjob design and its database reference relation.](https://drive.google.com/file/d/1ac503vDF_8zSTdFFuYwIlWSuyKFYg3ej/view?usp=drive_link)

