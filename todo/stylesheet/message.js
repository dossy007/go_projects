$(function () {
  $('span').click(function () {

    if ($(this).attr('class') == 'selected') { //出してる時
      // メニュー非表示
      $(this).removeClass('selected').next('ul').slideUp('fast');
    } else {
      // 表示しているメニューを閉じる
      $('span').removeClass('selected');
      $('ul').hide();

      // メニュー表示
      $(this).addClass('selected').next('ul').slideDown('fast');
    }
  });

  // マウスカーソルがメニュー上/メニュー外
  $('span,ul').hover(function () {
    over_flg = true;
  }, function () {
    over_flg = false; //メニュー外にいる時はfalse
  });

  // メニュー領域外をクリックしたらメニューを閉じる
  $('body').click(function () {
    if (over_flg == false) {
      $('span').removeClass('selected');
      $('ul').slideUp('fast');
    }
  });
});