# Cross-Site Request Forgery Attacks

No website is an island. Because your website has a public URI, other sites will frequently link to to it. More inbound links to your site means more traffic and better search engine rankings.

However, not everybody linking to your site has good intentions. An attacker can trick a user into clicking a malicious link that triggers underable or unexpected side effects. This is called CSRF.

## Anatomy of a CSRF Attack

Attackers usually launch CSRF attacks by exploiting websites that implement GET requests that change the state of a web server. A GET request is triggered when a victim clicks a link, allowing the attacker to craft misleading links into the target site that perform unexpected actions. GET requests are the only type of HTTP request that contain the entirely of the request's content in a URL, so they're uniquely vulerable to CSRF attacks.

```go
type Tweet struct {
	Title   string
	Content string
}

// Run ...
func Run(db *sqlx.DB) {
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		db.NamedExec("INSERT INTO tweets (:title, :content)", Tweet{
			Title:   fmt.Sprintf("title_%d", time.Now().UnixMilli()),
			Content: fmt.Sprintf("content_%d", time.Now().UnixMilli()),
		})
		w.Write([]byte("create success"))
	})

	http.ListenAndServe(":8080", nil)
}
```

If victim user click `http://localhost:8080/create`, then create new tweet.

Because they could use a single GET request to write a tweet, they constructed a malicious link that, when clicked, would post a new tweet containing an absene link that the first victim tweeted, they too were tricked into tweeting the same thing. 

The hacker tricked a handful of victims into clicking the malicious link, and those victims tweeted unexpected posts on the timelines.

### Mitigation 1: Follow REST Principles

To protect your users against CSRF attacks, make sure that your GET requests don't change the state of the server. Your website should use GET requests only to fetch web pages or other resources. You should perform actions that change server state.

Protecting your GET requests doesn't mean that there aren't vulerabilities in other types of requests, as you'll see with our second mitigation.

### Mitigation 2: Implement Anti-CSRF Cookies