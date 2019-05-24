# s3mem
s3mem is an in-memory S3API implementation. 
It doesn't require any server or external executable.
Useful for unit-testing as it doesn't require any external server.
It is a work in progess and only some APIs are implemented (Feel free to contribute)
A panic() is raised if you reach a non-implemented API.


## Example

    ```
    func TestListBucketsRequest(t *testing.T) {
        //Need to lock for testing as tests are running concurrently
        //and meanwhile another running test could change the stored buckets
        S3MemBuckets.Mux.Lock()
        defer S3MemBuckets.Mux.Unlock()

        //Adding bucket directly in mem to prepare the test.
        bucket0 := strings.ToLower(t.Name() + "0")
        bucket1 := strings.ToLower(t.Name() + "1")
        AddBucket(&s3.Bucket{Name: &bucket0})
        AddBucket(&s3.Bucket{Name: &bucket1})
        //Request a client
        client, err := NewClient()
        assert.NoError(t, err)
        assert.NotNil(t, client)
        //Create the request
        req := client.ListBucketsRequest(&s3.ListBucketsInput{})
        //Send the request
        listBucketsOutput, err := req.Send(context.Background())
        //Assert the result
        assert.NoError(t, err)
        assert.Equal(t, 2, len(listBucketsOutput.Buckets))
    }
    ```

## Limitation
- Pagination is not implemented, all items are returned.
- Only version "1" of object is available. 