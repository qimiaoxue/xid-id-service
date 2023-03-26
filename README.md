# Globally Unique ID Generator

Package xid is a globally unique id generator library, ready to safely be used directly in your server code.

Xid users the Mongo Object ID algorithm to generate globally unique ids with a different serialization (base64) to make it shorter when transported as a string: https://docs.mongodb.org/manual/reference/object-id/

- 4-byte value representing the seconds since the Unix epoch,
- 3-byte machine identifier
- 2-byte process id, and
- 3-byte counter, starting with a random value

The binary representation of the id is  compatible with Mongo 12 bytes Object IDs. The string representation is using base32 hex(w/o padding) for better space efficiency when stored in that form(20 bytes). The hex variant of base32 is used to retain the sortable property of the id.

Xid doesn't use base64 because case sensitivity and the 2 non alphanum chars may be an issue when transported as a string between various systems. Base36 wasn't retained either because 1/ it's not standard 2/the resulting size is not predictable (not bit aligned) and 2/ it would not remain sortable. To validate a base32 xid, expect a 20 chars long , all lowercase sequence of a to v letters and 0 to 9 numbers([0-9a-v]{20})

UUID are 16 bytes(128bits) and 36 chars as string representation. Twitter Snowflake ids are 8 bytes(64bits) but require machine/data-center configuration and/or central generator servers. Xid stands in between with 12 bytes(96 bits) and a more compact URL-safe string representation(20 chars). No configuration or central generator server is required so it can be used directly in server's code.

| Name  | Binary Size  | String Size  | Features       |
| :-------- | :------ | :----: |
| UUID | 16 bytes | 36 chars | configuration free, not sortable |
| shortuuid | 16 bytes | 22 chars | configuration free, not sortable |
| Snowflake | 8 bytes | up to 20 chars | needs machine/DC configuration, needs central server, sortable |
| MongoID | 12 bytes | 24 chars | configuration free, sortable |
| xid | 12 bytes | 20 chars | configuration free, sortable |

Features:
- Size: 12 bytes(96 bits), smaller than UUID, larger than snowflake
- Base32 hex encoded by default(20 chars when transported as printable string, still sortable)
- Non configured, you don't need set a unique machine and/or data center id
- K-ordered
- Embedded time with 1 second precision
- Unicity guaranteed for 16777216(24bits) unique ids per second and per host/process
- Lock-free(i.e.:unlike UUIDv1 and v2)

Best used with zerolog's RequestIDHandler.

Notes:
- Xid is dependent on the system time, a monotonic counter and so is not cryptographically secure. If unpredictability of IDs is important, you should use libraries that rely on cryptographically secure sources(like /dev/urandom on unix, crypto/rand in golang), if you want a truly random ID generator.

# Install
go get github.com/qimiaoxue/xid-id-service 

# Usage
guid := xid.New()
println(guid.String())
//Output: 9m4e2mr0ui3e8a215n4g

Get xid embeded info:

guid.Machine()
guid.Pid()
guid.Time()
guid.Counter()