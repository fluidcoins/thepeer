# thepeer
ThePeer SDK ( Go )

### Installation

```sh
go get github.com/fluidcoins/thepeer
```

### Usage

```go
  	c,err := New(WithAPISecret("API_KEY"))
	// check error
	// Also you can pass in your own *http.Client as follows
	// New(WithAPISecret("API_KEY"), WithHTTPClient(httpClient))
	// Although, you have to make sure the provided client is authenticated with your api key
	opts := &IndexUserOptions{
		Name:       randomdata.FullName(randomdata.RandomGender),
		Email:      randomdata.Email(),
		Identifier: randomdata.StringNumber(10, ""),
	}

	idxUser, err := c.IndexUser(opts)

	newIdxUser, err = c.UpdateUser(&UpdateUserOptions{
		Identifier: randomdata.StringNumber(10, ""),
		Reference:  idxUser.Reference,
	})
	require.NoError(t, err)
	require.NotEqual(t, idxUser.Identifier, newIdxUser.Identifier)

	c.DeIndexUser(&DeIndexUserOptions{
		UserReference: idxUser.Reference,
	})


```

### Status
- [x] Index user
- [x] De-index user
- [x] Update user
- [x] Fetch/Verify a receipt
- [x] Send receipt
