<?php include_once 'inc/header.php' ?>
<?php
	$posts = scandir('posts');
	foreach($posts as $post){
		if($post!=='.' && $post!=='..'){
			include_once 'posts/'.$post;
			echo '<hr/>';
			/*$post_link = str_replace('.html','',$post);
			$post_link = '<li><a href="'.$post_link.'/">'.strtoupper(str_replace('-',' ',$post_link)).'</a></li>';
			echo $post_link;*/
		}
	}
?>
<?php include_once 'inc/footer.php' ?>