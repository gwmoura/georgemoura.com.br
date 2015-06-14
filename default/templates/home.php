<div class="posts">
  <h4>Posts</h4>
  <ul>
  <?php
    $posts = scandir('../templates/posts', 1);
    foreach($posts as $post){
      if($post!=='.' && $post!=='..'){
        $post_link = substr($post, 9);
        $post_link = str_replace('.html','',$post_link);
        $post_link = '<li><a href="'.$post_link.'/">'.strtoupper(str_replace('-',' ',$post_link)).'</a></li>';
        echo $post_link;
      }
    }
  ?>
  </ul>
</div>
<div class="contatcts">
  <h4>Contatos</h4>
  <p>Tel: <a href="tel:07187886089">071 8788-6089</a></p>
  <p>E-mail: <a href="mailto:george@georgemoura.com.br">george@georgemoura.com.br</a></p>
  <p><a href="/feed/" title="Assinar Posts">RSS</a></p>
</div>