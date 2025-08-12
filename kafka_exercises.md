# 📌 Focused Kafka Integration: Team & Asset Changes

These exercises help integrate **Kafka** into the Team Management and Asset Sharing parts of your microservices system.

---

## 👥 Team Event Notification via Kafka

**Goal**: Emit events whenever team data changes—creation, member updates, or manager changes.

- **Topic**: `team.activity`
- **Events**:
  - `TEAM_CREATED`
  - `MEMBER_ADDED`
  - `MEMBER_REMOVED`
  - `MANAGER_ADDED`
  - `MANAGER_REMOVED`
- **Payload Example**:
  ```json
  {
    "eventType": "MEMBER_ADDED",
    "teamId": "uuid",
    "performedBy": "userId",
    "targetUserId": "userId",
    "timestamp": "ISO-8601"
  }
  ```
- **Consumer Suggestions**:
  - Log to database or ElasticSearch.
  - Send notifications to team members or admin UI.

---

## 🗂 Asset Change Stream via Kafka

**Goal**: Emit events when assets (folders/notes) are created, updated, deleted, or shared.

- **Topic**: `asset.changes`
- **Events**:
  - `FOLDER_CREATED`, `FOLDER_UPDATED`, `FOLDER_DELETED`
  - `NOTE_CREATED`, `NOTE_UPDATED`, `NOTE_DELETED`
  - `FOLDER_SHARED`, `FOLDER_UNSHARED`
  - `NOTE_SHARED`, `NOTE_UNSHARED`
- **Payload Example**:
  ```json
  {
    "eventType": "NOTE_UPDATED",
    "assetType": "note",
    "assetId": "uuid",
    "ownerId": "userId",
    "actionBy": "userId",
    "timestamp": "ISO-8601"
  }
  ```
- **Consumer Suggestions**:
  - Maintain audit logs.
  - Trigger cache invalidation or real-time UI updates.
  - Index assets into search system.

---

## 🔧 Technical Notes

- Use `confluent-kafka-go` or `segmentio/kafka-go` libraries for Go.
- Ensure all events include `eventType`, `timestamp`, and identifiers.
- Consider retry & error handling in Kafka consumers.