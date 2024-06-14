package schema

type Binary string

// Execute messages
//
// StoreObject store an object to the bucket and make the sender the owner of the object.
// The object is referenced by the hash of its content and this value is returned. If the
// object is already stored, an error is returned.
//
// The "pin" parameter specifies if the object should be pinned for the sender. In such
// case, the object cannot be removed (forget) from the storage.
//
// The "compression_algorithm" parameter specifies the algorithm for compressing the object
// before storing it in the storage, which is optional. If no algorithm is specified, the
// algorithm used is the first algorithm of the bucket configuration limits. Note that the
// chosen algorithm can save storage space, but it will increase CPU usage. Depending on the
// chosen compression algorithm and the achieved compression ratio, the gas cost of the
// operation will vary, either increasing or decreasing.
//
// ForgetObject first unpin the object from the bucket for the considered sender, then
// remove it from the storage if it is not pinned anymore. If the object is pinned for other
// senders, it is not removed from the storage and an error is returned. If the object is
// not pinned for the sender, this is a no-op.
//
// PinObject pins the object in the bucket for the considered sender. If the object is
// already pinned for the sender, this is a no-op. While an object is pinned, it cannot be
// removed from the storage.
//
// UnpinObject unpins the object in the bucket for the considered sender. If the object is
// not pinned for the sender, this is a no-op. The object can be removed from the storage if
// it is not pinned anymore.
type ExecuteMsg struct {
	StoreObject  *StoreObject  `json:"store_object,omitempty"`
	ForgetObject *ForgetObject `json:"forget_object,omitempty"`
	PinObject    *PinObject    `json:"pin_object,omitempty"`
	UnpinObject  *UnpinObject  `json:"unpin_object,omitempty"`
}

type ForgetObject struct {
	ID string `json:"id"`
}

type PinObject struct {
	ID string `json:"id"`
}

type StoreObject struct {
	// Specifies the compression algorithm to use when storing the object. If None, the first               
	// algorithm specified in the list of accepted compression algorithms of the bucket is used             
	// (see [BucketLimits::accepted_compression_algorithms]).                                               
	CompressionAlgorithm                                                                       *Passthrough `json:"compression_algorithm"`
	// The content of the object to store.                                                                  
	Data                                                                                       string       `json:"data"`
	// Specifies if the object should be pinned for the sender.                                             
	Pin                                                                                        bool         `json:"pin"`
}

type UnpinObject struct {
	ID string `json:"id"`
}

// Instantiate messages
type InstantiateMsg struct {
	// The name of the bucket. The name could not be empty or contains whitespaces. If name                          
	// contains whitespace, they will be removed.                                                                    
	Bucket                                                                                 string                    `json:"bucket"`
	// The configuration of the bucket.                                                                              
	Config                                                                                 *InstantiateMsgConfig     `json:"config,omitempty"`
	// The limits of the bucket.                                                                                     
	Limits                                                                                 *InstantiateMsgLimits     `json:"limits,omitempty"`
	// The configuration for paginated query.                                                                        
	Pagination                                                                             *InstantiateMsgPagination `json:"pagination,omitempty"`
}

// The configuration of the bucket.
//
// BucketConfig is the type of the configuration of a bucket.
//
// The configuration is set at the instantiation of the bucket, and is immutable and cannot
// be changed. The configuration is optional and if not set, the default configuration is
// used.
type InstantiateMsgConfig struct {
	// The acceptable compression algorithms for the objects in the bucket. If this parameter is               
	// not set (none or empty array), then all compression algorithms are accepted. If this                    
	// parameter is set, then only the compression algorithms in the array are accepted.                       
	//                                                                                                         
	// When an object is stored in the bucket without a specified compression algorithm, the                   
	// first algorithm in the array is used. Therefore, the order of the algorithms in the array               
	// is significant. Typically, the most efficient compression algorithm, such as the                        
	// NoCompression algorithm, should be placed first in the array.                                           
	//                                                                                                         
	// Any attempt to store an object using a different compression algorithm than the ones                    
	// specified here will fail.                                                                               
	AcceptedCompressionAlgorithms                                                               []Passthrough  `json:"accepted_compression_algorithms,omitempty"`
	// The algorithm used to hash the content of the objects to generate the id of the objects.                
	// The algorithm is optional and if not set, the default algorithm is used.                                
	//                                                                                                         
	// The default algorithm is Sha256 if not set.                                                             
	HashAlgorithm                                                                               *HashAlgorithm `json:"hash_algorithm,omitempty"`
}

// The limits of the bucket.
//
// BucketLimits is the type of the limits of a bucket.
//
// The limits are optional and if not set, there is no limit.
type InstantiateMsgLimits struct {
	// The maximum number of pins in the bucket for an object.        
	MaxObjectPins                                             *string `json:"max_object_pins"`
	// The maximum size of the objects in the bucket.                 
	MaxObjectSize                                             *string `json:"max_object_size"`
	// The maximum number of objects in the bucket.                   
	MaxObjects                                                *string `json:"max_objects"`
	// The maximum total size of the objects in the bucket.           
	MaxTotalSize                                              *string `json:"max_total_size"`
}

// The configuration for paginated query.
//
// PaginationConfig is the type carrying configuration for paginated queries.
//
// The fields are optional and if not set, there is a default configuration.
type InstantiateMsgPagination struct {
	// The default number of elements in a page.                                      
	//                                                                                
	// Shall be less or equal than `max_page_size`. Default to '10' if not set.       
	DefaultPageSize                                                            *int64 `json:"default_page_size,omitempty"`
	// The maximum elements a page can contain.                                       
	//                                                                                
	// Shall be less than `u32::MAX - 1`. Default to '30' if not set.                 
	MaxPageSize                                                                *int64 `json:"max_page_size,omitempty"`
}

// Query messages
//
// Bucket returns the bucket information.
//
// Object returns the object information with the given id.
//
// Objects returns the list of objects in the bucket with support for pagination.
//
// ObjectData returns the content of the object with the given id.
//
// ObjectPins returns the list of addresses that pinned the object with the given id with
// support for pagination.
type QueryMsg struct {
	Bucket     *Bucket     `json:"bucket,omitempty"`
	Object     *Object     `json:"object,omitempty"`
	Objects    *Objects    `json:"objects,omitempty"`
	ObjectData *ObjectData `json:"object_data,omitempty"`
	ObjectPins *ObjectPins `json:"object_pins,omitempty"`
}

type Bucket struct {
}

type Object struct {
	// The id of the object to get.       
	ID                             string `json:"id"`
}

type ObjectData struct {
	// The id of the object to get.       
	ID                             string `json:"id"`
}

type ObjectPins struct {
	// The point in the sequence to start returning pins.        
	After                                                *string `json:"after"`
	// The number of pins to return.                             
	First                                                *int64  `json:"first"`
	// The id of the object to get the pins for.                 
	ID                                                   string  `json:"id"`
}

type Objects struct {
	// The owner of the objects to get.                             
	Address                                                 *string `json:"address"`
	// The point in the sequence to start returning objects.        
	After                                                   *string `json:"after"`
	// The number of objects to return.                             
	First                                                   *int64  `json:"first"`
}

// BucketResponse is the response of the Bucket query.
type BucketResponse struct {
	// The configuration of the bucket.                               
	Config                                   BucketResponseConfig     `json:"config"`
	// The limits of the bucket.                                      
	Limits                                   BucketResponseLimits     `json:"limits"`
	// The name of the bucket.                                        
	Name                                     string                   `json:"name"`
	// The configuration for paginated query.                         
	Pagination                               BucketResponsePagination `json:"pagination"`
}

// The configuration of the bucket.
//
// BucketConfig is the type of the configuration of a bucket.
//
// The configuration is set at the instantiation of the bucket, and is immutable and cannot
// be changed. The configuration is optional and if not set, the default configuration is
// used.
type BucketResponseConfig struct {
	// The acceptable compression algorithms for the objects in the bucket. If this parameter is               
	// not set (none or empty array), then all compression algorithms are accepted. If this                    
	// parameter is set, then only the compression algorithms in the array are accepted.                       
	//                                                                                                         
	// When an object is stored in the bucket without a specified compression algorithm, the                   
	// first algorithm in the array is used. Therefore, the order of the algorithms in the array               
	// is significant. Typically, the most efficient compression algorithm, such as the                        
	// NoCompression algorithm, should be placed first in the array.                                           
	//                                                                                                         
	// Any attempt to store an object using a different compression algorithm than the ones                    
	// specified here will fail.                                                                               
	AcceptedCompressionAlgorithms                                                               []Passthrough  `json:"accepted_compression_algorithms,omitempty"`
	// The algorithm used to hash the content of the objects to generate the id of the objects.                
	// The algorithm is optional and if not set, the default algorithm is used.                                
	//                                                                                                         
	// The default algorithm is Sha256 if not set.                                                             
	HashAlgorithm                                                                               *HashAlgorithm `json:"hash_algorithm,omitempty"`
}

// The limits of the bucket.
//
// BucketLimits is the type of the limits of a bucket.
//
// The limits are optional and if not set, there is no limit.
type BucketResponseLimits struct {
	// The maximum number of pins in the bucket for an object.        
	MaxObjectPins                                             *string `json:"max_object_pins"`
	// The maximum size of the objects in the bucket.                 
	MaxObjectSize                                             *string `json:"max_object_size"`
	// The maximum number of objects in the bucket.                   
	MaxObjects                                                *string `json:"max_objects"`
	// The maximum total size of the objects in the bucket.           
	MaxTotalSize                                              *string `json:"max_total_size"`
}

// The configuration for paginated query.
//
// PaginationConfig is the type carrying configuration for paginated queries.
//
// The fields are optional and if not set, there is a default configuration.
type BucketResponsePagination struct {
	// The default number of elements in a page.                                      
	//                                                                                
	// Shall be less or equal than `max_page_size`. Default to '10' if not set.       
	DefaultPageSize                                                            *int64 `json:"default_page_size,omitempty"`
	// The maximum elements a page can contain.                                       
	//                                                                                
	// Shall be less than `u32::MAX - 1`. Default to '30' if not set.                 
	MaxPageSize                                                                *int64 `json:"max_page_size,omitempty"`
}

// ObjectResponse is the response of the Object query.
type ObjectResponse struct {
	// The size of the object when compressed. If the object is not compressed, the value is the            
	// same as `size`.                                                                                      
	CompressedSize                                                                              string      `json:"compressed_size"`
	// The compression algorithm used to compress the content of the object.                                
	CompressionAlgorithm                                                                        Passthrough `json:"compression_algorithm"`
	// The id of the object.                                                                                
	ID                                                                                          string      `json:"id"`
	// Tells if the object is pinned by at least one address.                                               
	IsPinned                                                                                    bool        `json:"is_pinned"`
	// The owner of the object.                                                                             
	Owner                                                                                       string      `json:"owner"`
	// The size of the object.                                                                              
	Size                                                                                        string      `json:"size"`
}

// ObjectPinsResponse is the response of the GetObjectPins query.
type ObjectPinsResponse struct {
	// The list of addresses that pinned the object.                           
	Data                                            []string                   `json:"data"`
	// The page information.                                                   
	PageInfo                                        ObjectPinsResponsePageInfo `json:"page_info"`
}

// The page information.
//
// PageInfo is the page information returned for paginated queries.
type ObjectPinsResponsePageInfo struct {
	// The cursor to the next page.         
	Cursor                           string `json:"cursor"`
	// Tells if there is a next page.       
	HasNextPage                      bool   `json:"has_next_page"`
}

// ObjectsResponse is the response of the Objects query.
type ObjectsResponse struct {
	// The list of objects in the bucket.                        
	Data                                 []DatumElement          `json:"data"`
	// The page information.                                     
	PageInfo                             ObjectsResponsePageInfo `json:"page_info"`
}

// ObjectResponse is the response of the Object query.
type DatumElement struct {
	// The size of the object when compressed. If the object is not compressed, the value is the            
	// same as `size`.                                                                                      
	CompressedSize                                                                              string      `json:"compressed_size"`
	// The compression algorithm used to compress the content of the object.                                
	CompressionAlgorithm                                                                        Passthrough `json:"compression_algorithm"`
	// The id of the object.                                                                                
	ID                                                                                          string      `json:"id"`
	// Tells if the object is pinned by at least one address.                                               
	IsPinned                                                                                    bool        `json:"is_pinned"`
	// The owner of the object.                                                                             
	Owner                                                                                       string      `json:"owner"`
	// The size of the object.                                                                              
	Size                                                                                        string      `json:"size"`
}

// The page information.
//
// PageInfo is the page information returned for paginated queries.
type ObjectsResponsePageInfo struct {
	// The cursor to the next page.         
	Cursor                           string `json:"cursor"`
	// Tells if there is a next page.       
	HasNextPage                      bool   `json:"has_next_page"`
}

// Represents no compression algorithm. The object is stored as is without any compression.
//
// Represents the Snappy algorithm. Snappy is a compression/decompression algorithm that
// does not aim for maximum compression. Instead, it aims for very high speeds and
// reasonable compression.
//
// See [the snappy web page](https://google.github.io/snappy/) for more information.
//
// Represents the LZMA algorithm. LZMA is a lossless data compression/decompression
// algorithm that features a high compression ratio and a variable compression-dictionary
// size up to 4 GB.
//
// See [the LZMA wiki
// page](https://en.wikipedia.org/wiki/Lempel%E2%80%93Ziv%E2%80%93Markov_chain_algorithm)
// for more information.
//
// CompressionAlgorithm is an enumeration that defines the different compression algorithms
// supported for compressing the content of objects. The compression algorithm specified
// here are relevant algorithms for compressing data on-chain, which means that they are
// fast to compress and decompress, and have a low computational cost.
//
// The order of the compression algorithms is based on their estimated computational cost
// (quite opinionated) during both compression and decompression, ranging from the lowest to
// the highest. This particular order is utilized to establish the default compression
// algorithm for storing an object.
//
// The compression algorithm used to compress the content of the object.
type Passthrough string

const (
	Lzma                   Passthrough = "lzma"
	PassthroughPassthrough Passthrough = "passthrough"
	Snappy                 Passthrough = "snappy"
)

// The algorithm used to hash the content of the objects to generate the id of the objects.
// The algorithm is optional and if not set, the default algorithm is used.
//
// The default algorithm is Sha256 if not set.
//
// HashAlgorithm is an enumeration that defines the different hash algorithms supported for
// hashing the content of objects.
//
// Represents the MD5 algorithm. MD5 is a widely used cryptographic hash function that
// produces a 128-bit hash value. The computational cost of MD5 is relatively low compared
// to other hash functions, but its short hash length makes it easier to find hash
// collisions. It is now considered insecure for cryptographic purposes, but can still used
// in non-security contexts.
//
// MD5 hashes are stored on-chain as 32 hexadecimal characters.
//
// See [the MD5 Wikipedia page](https://en.wikipedia.org/wiki/MD5) for more information.
//
// Represents the SHA-224 algorithm. SHA-224 is a variant of the SHA-2 family of hash
// functions that produces a 224-bit hash value. It is similar to SHA-256, but with a
// shorter output size. The computational cost of SHA-224 is moderate, and its relatively
// short hash length makes it easier to store and transmit.
//
// SHA-224 hashes are stored on-chain as 56 hexadecimal characters.
//
// See [the SHA-2 Wikipedia page](https://en.wikipedia.org/wiki/SHA-2) for more
// information.
//
// Represents the SHA-256 algorithm. SHA-256 is a member of the SHA-2 family of hash
// functions that produces a 256-bit hash value. It is widely used in cryptography and other
// security-related applications. The computational cost of SHA-256 is moderate, and its
// hash length strikes a good balance between security and convenience.
//
// SHA-256 hashes are stored on-chain as 64 hexadecimal characters.
//
// See [the SHA-2 Wikipedia page](https://en.wikipedia.org/wiki/SHA-2) for more
// information.
//
// Represents the SHA-384 algorithm. SHA-384 is a variant of the SHA-2 family of hash
// functions that produces a 384-bit hash value. It is similar to SHA-512, but with a
// shorter output size. The computational cost of SHA-384 is relatively high, but its longer
// hash length provides better security against hash collisions.
//
// SHA-384 hashes are stored on-chain as 96 hexadecimal characters.
//
// See [the SHA-2 Wikipedia page](https://en.wikipedia.org/wiki/SHA-2) for more
// information.
//
// Represents the SHA-512 algorithm. SHA-512 is a member of the SHA-2 family of hash
// functions that produces a 512-bit hash value. It is widely used in cryptography and other
// security-related applications. The computational cost of SHA-512 is relatively high, but
// its longer hash length provides better security against hash collisions.
//
// SHA-512 hashes are stored on-chain as 128 hexadecimal characters.
//
// See [the SHA-2 Wikipedia page](https://en.wikipedia.org/wiki/SHA-2) for more information.
type HashAlgorithm string

const (
	MD5    HashAlgorithm = "m_d5"
	Sha224 HashAlgorithm = "sha224"
	Sha256 HashAlgorithm = "sha256"
	Sha384 HashAlgorithm = "sha384"
	Sha512 HashAlgorithm = "sha512"
)
