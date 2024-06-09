# GraphQL API Documentation

This API allows you to perform various operations related to managing videos. It supports the following operations:

## 1. Upload a Video

To upload a new video, first, the frontend user should upload the video file to a third-party service like Cloudinary. Once the video is uploaded, Cloudinary will provide the necessary URLs and metadata for the video file and thumbnail.

After obtaining the video file URL, thumbnail URL, and other metadata from Cloudinary, the frontend user should send a `mutation` request with the `uploadVideo` operation and provide the required input data, including the title and description entered by the user on the frontend.

**Mutation:**

```graphql
mutation {
  uploadVideo(
    input: {
      title: "Video Title" # Entered by the user on the frontend
      description: "Video Description" # Entered by the user on the frontend
      duration: 8 # Obtained from Cloudinary metadata
      thumbnailUrl: "https://res.cloudinary.com/dyxrrljec/video/upload/v1717856605/ynvzpktnc2q8icemdlki.mp4" # Obtained from Cloudinary
      file: "https://res.cloudinary.com/dyxrrljec/video/upload/v1717856605/ynvzpktnc2q8icemdlki.mp4" # Obtained from Cloudinary
    }
  ) {
    id
    title
    description
    duration
    thumbnailUrl
    videoUrl
    createdAt
    updatedAt
  }
}
```

**Input Fields:**

- `title` (String): The title of the video, entered by the user on the frontend.
- `description` (String): The description of the video, entered by the user on the frontend.
- `duration` (Int): The duration of the video in seconds, obtained from Cloudinary metadata.
- `thumbnailUrl` (String): The URL of the video thumbnail, obtained from Cloudinary.
- `file` (String): The URL of the video file, obtained from Cloudinary.

**Response:**

```json
{
  "data": {
    "uploadVideo": {
      "id": "d83e5438-987a-40fb-9c78-8bafbc5fb92e",
      "title": "Video Title",
      "description": "Video Description",
      "duration": 8,
      "thumbnailUrl": "https://res.cloudinary.com/dyxrrljec/video/upload/v1717856605/ynvzpktnc2q8icemdlki.mp4",
      "videoUrl": "https://res.cloudinary.com/dyxrrljec/video/upload/v1717856605/ynvzpktnc2q8icemdlki.mp4",
      "createdAt": "2024-06-09T10:48:30.380648417Z",
      "updatedAt": "2024-06-09T10:48:30.380648417Z"
    }
  }
}
```

## 2. Get All Videos

To retrieve a list of all uploaded videos, send a `query` request with the `videos` operation.

**Query:**

```graphql
{
  videos {
    id
    title
    description
    duration
    thumbnailUrl
    videoUrl
    createdAt
    updatedAt
  }
}
```

**Response:**

```json
{
  "data": {
    "videos": [
      {
        "id": "5a36bae9-afe9-4ffc-9dc6-8f3dbda35032",
        "title": "Video Title",
        "description": "Video Description",
        "duration": 8,
        "thumbnailUrl": "https://res.cloudinary.com/dyxrrljec/video/upload/v1717856605/ynvzpktnc2q8icemdlki.mp4",
        "videoUrl": "https://res.cloudinary.com/dyxrrljec/video/upload/v1717856605/ynvzpktnc2q8icemdlki.mp4",
        "createdAt": "2024-06-08T17:24:04.541875+03:00",
        "updatedAt": "2024-06-08T17:24:04.541875+03:00"
      }
    ]
  }
}
```

## 3. Get a Specific Video

To retrieve details of a specific video, send a `query` request with the `video` operation and provide the `id` of the video you want to retrieve.

**Query:**

```graphql
{
  video(id: "string") {
    id
    title
    description
    duration
    thumbnailUrl
    videoUrl
    createdAt
    updatedAt
  }
}
```

**Input Fields:**

- `id` (String): The ID of the video you want to retrieve.

**Response:**

```json
{
  "data": {
    "video": {
      "id": "5a36bae9-afe9-4ffc-9dc6-8f3dbda35032",
      "title": "Video Title",
      "description": "Video Description",
      "duration": 8,
      "thumbnailUrl": "https://res.cloudinary.com/dyxrrljec/video/upload/v1717856605/ynvzpktnc2q8icemdlki.mp4",
      "videoUrl": "https://res.cloudinary.com/dyxrrljec/video/upload/v1717856605/ynvzpktnc2q8icemdlki.mp4",
      "createdAt": "2024-06-08T17:24:04.541875+03:00",
      "updatedAt": "2024-06-08T17:24:04.541875+03:00"
    }
  }
}
```

In summary, to upload a new video, the frontend user should first upload the video file to a third-party service like Cloudinary, obtain the necessary URLs and metadata, and then send a mutation request to the GraphQL API with the video details, including the title and description entered by the user on the frontend.

Note: Make sure to replace the placeholder values (`"string"`, `"Test title"`, `"Test description"`, etc.) with the actual values or variables used in your API.
