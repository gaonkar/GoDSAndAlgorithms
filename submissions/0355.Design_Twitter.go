/*
Design a simplified version of Twitter where users can post tweets, follow/unfollow another user and is able to see the 10 most recent tweets in the user's news feed. Your design should support the following methods:

postTweet(userId, tweetId): Compose a new tweet.
getNewsFeed(userId): Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent.
follow(followerId, followeeId): Follower follows a followee.
unfollow(followerId, followeeId): Follower unfollows a followee.
Example:

Twitter twitter = new Twitter();

// User 1 posts a new tweet (id = 5).
twitter.postTweet(1, 5);

// User 1's news feed should return a list with 1 tweet id -> [5].
twitter.getNewsFeed(1);

// User 1 follows user 2.
twitter.follow(1, 2);

// User 2 posts a new tweet (id = 6).
twitter.postTweet(2, 6);

// User 1's news feed should return a list with 2 tweet ids -> [6, 5].
// Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
twitter.getNewsFeed(1);

// User 1 unfollows user 2.
twitter.unfollow(1, 2);

// User 1's news feed should return a list with 1 tweet id -> [5],
// since user 1 is no longer following user 2.
twitter.getNewsFeed(1);
*/
type P  struct{
    Val int
    tid int
}
type PHeap []P

func (h PHeap) Len() int           { return len(h) }
func (h PHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h PHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PHeap) Push(x interface{}) {
        // Push and Pop use pointer receivers because they modify the slice's length,
        // not just its contents.
        *h = append(*h, x.(P))
}

func (h *PHeap) Pop() interface{} {
        old := *h
        n := len(old)
        x := old[n-1]
        *h = old[0 : n-1]
        return x
}


type Tweet struct {
    gid int
    tid int
}

type Twitter struct {
    gtid    int
    M       int     //max tweets
    umap    map[int] []Tweet
    follow  map[int] map[int] bool
}


/** Initialize your data structure here. */
func Constructor() Twitter {
    var T Twitter
    T.umap =make(map[int] []Tweet)
    T.follow = make(map[int] map[int] bool)
    T.M = 10
    return T
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
    var w Tweet
    _, ok := this.umap[userId]
    w.gid = this.gtid
    w.tid = tweetId
    if !ok {
        this.umap[userId] = make([]Tweet, 0)
    }
    this.umap[userId]  = append(this.umap[userId], w)
    this.gtid++
    this.Follow(userId, userId)
}


func Reverse(r []int) {
    n := len(r)-1
    for i := 0; i < n;i++ {
        r[i], r[n] = r[n], r[i]
        n--
    }
}


func GetTopTweets(maxTweets int, umap map[int][]Tweet, ulist map[int] bool) (r []int) {
    H := &PHeap{}
    heap.Init(H)
    for k,_ := range(ulist) {
        for _,t := range(umap[k]) {
            heap.Push(H, P{t.gid, t.tid})
            if H.Len() > maxTweets {
                heap.Pop(H)
            }
        }
    }
    for H.Len() > 0 {
        x := heap.Pop(H).(P)
        r = append(r, x.tid)
    }
    Reverse(r)
    return r
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {

    ulist, ok := this.follow[userId]
    if !ok {
        return nil
    }

    return GetTopTweets(this.M, this.umap, ulist)
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
    _, ok := this.follow[followerId]
    if !ok {
        this.follow[followerId] = make(map[int] bool)
    }
    this.follow[followerId][followeeId] = true
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if followerId == followeeId {return }
    _, ok := this.follow[followerId]
    if !ok {
        return
    }
    _, ok  = this.follow[followerId][followeeId]
    if (!ok) {
        return
    }
    delete(this.follow[followerId], followeeId)

}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
