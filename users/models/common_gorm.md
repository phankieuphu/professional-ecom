
## ✅ Common GORM Tags with Examples

| Tag | Description | Example |
|-----|-------------|---------|
| `primaryKey` | Marks a field as the primary key | `ID uint \`gorm:"primaryKey"\`` |
| `autoIncrement` | Auto-increment a numeric field | `Counter int \`gorm:"autoIncrement"\`` |
| `not null` | Prevents null values | `Name string \`gorm:"not null"\`` |
| `unique` or `uniqueIndex` | Adds a UNIQUE constraint | `Email string \`gorm:"uniqueIndex"\`` |
| `default` | Sets a default value | `Age int \`gorm:"default:18"\`` |
| `size` | Sets size for string fields (VARCHAR) | `Username string \`gorm:"size:50"\`` |
| `column` | Sets a custom DB column name | `FullName string \`gorm:"column:full_name"\`` |
| `index` | Creates an index on this field | `Country string \`gorm:"index"\`` |
| `type` | Manually sets DB column type | `Metadata string \`gorm:"type:json"\`` |
| `<-:create` | Field is writeable only on create | `CreatedBy string \`gorm:"<-:create"\`` |
| `->` | Read-only field (ignored on create/update) | `Score int \`gorm:"->"\`` |
| `-` | Ignored by GORM (not mapped to DB) | `TempData string \`gorm:"-"\`` |
| `embedded` / `embeddedPrefix` | Embeds struct fields | See below for example |
| `foreignKey`, `references` | Used in relations (FK setup) | See relationship examples |

---

