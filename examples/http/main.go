package main

import (
	"context"
	"fmt"

	"lurkers.club/lemmy-api/lemmy"
	"lurkers.club/lemmy-api/types"
)

func main() {
	ctx := context.Background()

	c, err := lemmy.New("https://voyager.lemmy.ml/")
	if err != nil {
		fmt.Print(err)
	}

	// For each I can retrieve max 50 at a time, need to use paging for all of them
	// e.g. max 50 communities, 50 posts, 50 comments

	getSite(c, ctx)

	printCommunitiesPostsAndComments(c, ctx)
}

func getSite(c *lemmy.Client, ctx context.Context) {
	if response, err := c.Site(ctx, types.GetSite{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("Site: %s, Total communities: %d",
			response.SiteView.Site.Name, response.SiteView.Counts.Communities))
	}
}

func printCommunitiesPostsAndComments(c *lemmy.Client, ctx context.Context) {
	if response, err := c.Communities(ctx, types.ListCommunities{Type: types.NewOptional(types.ListingTypeLocal),
		Limit: types.NewOptional(int64(50))}); err != nil {

			fmt.Print(err)
	} else {
		for _, communityView := range response.Communities {
			fmt.Println(fmt.Sprintf("Community: %s, Total posts: %d",
				communityView.Community.Name, communityView.Counts.Posts))
			printPostsAndComments(c, ctx, communityView.Community.ID)
			fmt.Println("-----------------------------------------------------")
		}
	}
}

func printPostsAndComments(c *lemmy.Client, ctx context.Context, communityId int) {
	if response, err := c.Posts(ctx, types.GetPosts{Type: types.NewOptional(types.ListingTypeLocal),
		CommunityID: types.NewOptional(communityId),
		Limit: types.NewOptional(int64(50))}); err != nil {

			fmt.Print(err)
	} else {
		for _, postView := range response.Posts {
			fmt.Println(fmt.Sprintf("Post: %s, Total comments: %d",
				postView.Post.Name, postView.Counts.Comments))

			printComments(c, ctx, postView)
		}
	}
}

func printComments(c *lemmy.Client, ctx context.Context, postView types.PostView) {
	// Need to figure how MaxDepth works
	if response, err := c.Comments(ctx, types.GetComments{PostID: types.NewOptional(postView.Post.ID),
		Limit: types.NewOptional(int64(50))}); err != nil {

			fmt.Println(err)
	} else {
		for _, commentView := range response.Comments {
			// How do I tell which is the parent comment?
			fmt.Println(fmt.Sprintf("Comment: %s, Path: %s, Total replies: %d",
				commentView.Comment.Content, commentView.Comment.Path, commentView.Counts.ChildCount))
		}
	}
}

// export default function buildCommentsTree(
//   comments: CommentView[],
//   parentComment: boolean
// ): CommentNodeI[] {
//   const map = new Map<number, CommentNodeI>();
//   const depthOffset = !parentComment
//     ? 0
//     : getDepthFromComment(comments[0].comment) ?? 0;

//   for (const comment_view of comments) {
//     const depthI = getDepthFromComment(comment_view.comment) ?? 0;
//     const depth = depthI ? depthI - depthOffset : 0;
//     const node: CommentNodeI = {
//       comment_view,
//       children: [],
//       depth,
//     };
//     map.set(comment_view.comment.id, { ...node });
//   }

//   const tree: CommentNodeI[] = [];

//   // if its a parent comment fetch, then push the first comment to the top node.
//   if (parentComment) {
//     const cNode = map.get(comments[0].comment.id);
//     if (cNode) {
//       tree.push(cNode);
//     }
//   }

//   for (const comment_view of comments) {
//     const child = map.get(comment_view.comment.id);
//     if (child) {
//       const parent_id = getCommentParentId(comment_view.comment);
//       if (parent_id) {
//         const parent = map.get(parent_id);
//         // Necessary because blocked comment might not exist
//         if (parent) {
//           parent.children.push(child);
//         }
//       } else {
//         if (!parentComment) {
//           tree.push(child);
//         }
//       }
//     }
//   }

//   return tree;
// }

// export interface CommentNodeI {
//   comment_view: CommentView;
//   children: Array<CommentNodeI>;
//   depth: number;
// }


// export default function getDepthFromComment(
//   comment?: Comment
// ): number | undefined {
//   const len = comment?.path.split(".").length;
//   return len ? len - 2 : undefined;
// }

// export default function getCommentParentId(
//   comment?: Comment
// ): number | undefined {
//   const split = comment?.path.split(".");
//   // remove the 0
//   split?.shift();

//   return split && split.length > 1
//     ? Number(split.at(split.length - 2))
//     : undefined;
// }
