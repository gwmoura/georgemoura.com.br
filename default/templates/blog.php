<?php
	$posts = scandir('../templates/posts', 1);
	foreach($posts as $post){
		if($post!=='.' && $post!=='..'){
			include_once '../templates/posts/'.$post;
			echo '<hr/>';
		}
	}
?>