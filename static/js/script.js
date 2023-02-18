$(function() {
    loadPosts();

    $('#post-form').submit(function(e) {
        e.preventDefault();
        var title = $('#title').val();
        var content = $('#content').val();
        var author = $('#author').val();
        var create_at = $('#create_at').val();

        $.ajax({
            type: 'POST',
            url: '/posts',
            data: JSON.stringify({title: title, content: content, author: author, create_at: create_at}),
            contentType: 'application/json',
            success: function(response) {
                loadPosts();
                $('#title').val('');
                $('#content').val('');
                $('#author').val('');
                $('#create_at').val('');
            }
        });
    });
});

function loadPosts() {
    $.get('/posts', function(response) {
        var posts = response.data;
        var $posts = $('#posts');

        $posts.empty();

        console.log(posts)
        $.each(posts, function(i, post) {
            var $post = $('<div>').addClass('post');
            var $title = $('<h2>').text(post.Title);
            var $author = $('<div>').text(post.Author + "-" + post.create_at);
            var $content = $('<div>').text(post.Content);

            $post.append($title).append($author).append($('<p>')).append($content);
            $posts.append($post);
        });
    });
}